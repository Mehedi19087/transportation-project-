package dailysidecash

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Service interface {
	Create(req *CreateDailySideCashDTO, productID uint) (*DailySideCash, error)
	Update(id uint, req *UpdateDailySideCashDTO) error
	Get(id uint) (*DailySideCash, error)
	GetAll(productID uint, page, pageSize int) ([]DailySideCash, int64, error)
	Delete(id uint) error
	GetByDate(productID uint, date string) (*DailySideCash, error)
}

type service struct {
	repo DailySideCashRepo
}

func NewService(repo DailySideCashRepo) Service {
	return &service{repo: repo}
}

func (s *service) Create(req *CreateDailySideCashDTO, productID uint) (*DailySideCash, error) {
	if req.Date == "" {
		return nil, errors.New("date is required")
	}
	d, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format (expected YYYY-MM-DD): %w", err)
	}

	// Fetch last available record before the current date
	var carryOver float64 = 0
	prevRecord, err := s.repo.GetLastRecordBeforeDate(productID, d)
	if err != nil {
		// Log error if needed, but treat as 0 carryOver if strictly database error, 
        // though implementation of GetLastRecordBeforeDate returns nil, nil for not found.
		carryOver = 0
	} else if prevRecord != nil {
		carryOver = prevRecord.RemainingBalance
	}

	// Total Cash = Carry Over from yesterday + Any new Cash added today (req.Cash)
	totalCash := carryOver + req.Cash

	// Calculate Remaining Balance
	remainingBalance := totalCash
	if req.TripCost > 0 || req.OtherCost > 0 {
		remainingBalance = totalCash - (req.TripCost + req.OtherCost)
	}

	withoutremaining := req.Cash

	record := &DailySideCash{
		Date:             d,
		ProductID:        productID,
		Cash:             totalCash,
		RemainingBalance: remainingBalance,
		TripCost:         req.TripCost,
		OtherCost:        req.OtherCost,
		OtherCostDetails: req.OtherCostDetails,
		WithoutRemaining: withoutremaining,
	}

	if err := s.repo.Create(record); err != nil {
		return nil, fmt.Errorf("create daily side cash: %w", err)
	}
	return record, nil
}

func (s *service) Update(id uint, req *UpdateDailySideCashDTO) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("daily side cash not found")
		}
		return err
	}

	if req.Date != nil {
		d, parseErr := time.Parse("2006-01-02", *req.Date)
		if parseErr != nil {
			return fmt.Errorf("invalid date format (expected YYYY-MM-DD): %w", parseErr)
		}
		existing.Date = d
	}
	if req.Cash != nil {
		existing.Cash += *req.Cash
	}
	if req.TripCost != nil {
		existing.TripCost += *req.TripCost
	}
	if req.OtherCost != nil {
		existing.OtherCost += *req.OtherCost
	}
	if req.OtherCostDetails != nil {
		existing.OtherCostDetails = *req.OtherCostDetails
	}

	oldRemaining := existing.RemainingBalance

	existing.RemainingBalance = existing.Cash - (existing.TripCost + existing.OtherCost)
	if req.Cash != nil {
		existing.WithoutRemaining = existing.WithoutRemaining + *req.Cash 
	}

	if err := s.repo.Update(existing); err != nil {
		return fmt.Errorf("update daily side cash: %w", err)
	}

	newRemaining := existing.RemainingBalance
	if oldRemaining != newRemaining {
		diff := newRemaining - oldRemaining
		
		nextRecord, err := s.repo.GetFirstRecordAfterDate(existing.ProductID, existing.Date)
		if err == nil && nextRecord != nil {
			nextRecord.Cash += diff
			nextRecord.RemainingBalance += diff
			if err := s.repo.Update(nextRecord); err != nil {
				return fmt.Errorf("failed to update next record: %w", err)
			}
		}
	}

	return nil
}
func (s *service) Get(id uint) (*DailySideCash, error) {
	rec, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("daily side cash not found")
		}
		return nil, err
	}
	return rec, nil
}

func (s *service) GetAll(productID uint, page, pageSize int) ([]DailySideCash, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	records, total, err := s.repo.GetAll(productID, offset, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("list daily side cash: %w", err)
	}
	return records, total, nil
}

func (s *service) Delete(id uint) error {
	// 1. Get the record before deleting to know its Date and Balance
	record, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("get daily side cash for delete: %w", err)
	}

	// 2. Calculate the difference for future records.
	// We need to know what the 'carryOver' (start balance) for the NEXT record will be
	// now that this record is gone. That 'carryOver' comes from the PREVIOUS record.
	var prevBalance float64 = 0
	prevRecord, err := s.repo.GetLastRecordBeforeDate(record.ProductID, record.Date)
	if err == nil && prevRecord != nil {
		prevBalance = prevRecord.RemainingBalance
	}
	
	// The future records currently have `record.RemainingBalance` factored in.
	// They SHOULD have `prevBalance` factored in instead.
	// Diff = New_Baseline - Old_Baseline
	diff := prevBalance - record.RemainingBalance

	// 3. Delete the record
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("delete daily side cash: %w", err)
	}

	// 4. Update all future records with the difference
	if err := s.repo.UpdateFutureBalances(record.ProductID, record.Date, diff); err != nil {
		return fmt.Errorf("failed to update future balances after delete: %w", err)
	}

	return nil
}

func (s *service) GetByDate(productID uint, date string) (*DailySideCash, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format (expected YYYY-MM-DD): %w", err)
	}
	rec, err := s.repo.GetByDate(productID, d)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("daily side cash not found")
		}
		return nil, err
	}
	return rec, nil
}
