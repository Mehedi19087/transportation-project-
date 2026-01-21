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

	// Calculate previous date
	prevDate := d.AddDate(0, 0, -1)
	//prevDateStr := prevDate.Format("2006-01-02")

	// Fetch previous day's record
	var carryOver float64 = 0
	prevRecord, err := s.repo.GetByDate(productID, prevDate)
	if err != nil {
		// If it's a real DB error (not just "not found"), return it so we can see what's wrong
		if !errors.Is(err, gorm.ErrRecordNotFound) && err.Error() != "daily side cash record not found for the given date and product" {
			return nil, fmt.Errorf("failed to fetch previous day record: %w", err)
		}
		// If it is just not found, we start with 0 (normal for first day)
		carryOver = 0
	} else if prevRecord != nil {
		carryOver = prevRecord.RemainingBalance
	}

	// Total Cash = Carry Over from yesterday + Any new Cash added today (req.Cash)
	totalCash := carryOver + req.Cash

	// Calculate Remaining Balance
	remainingBalance := totalCash - (req.TripCost + req.OtherCost)

	withoutremaining:= req.Cash 

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
	if req.RemainingBalance != nil {
		existing.RemainingBalance = *req.RemainingBalance
	}
	if req.TripCost != nil {
		existing.TripCost = *req.TripCost
	}
	if req.OtherCost != nil {
		existing.OtherCost = *req.OtherCost
	}
	if req.OtherCostDetails != nil {
		existing.OtherCostDetails = *req.OtherCostDetails
	}
	existing.WithoutRemaining = existing.Cash

	if err := s.repo.Update(existing); err != nil {
		return fmt.Errorf("update daily side cash: %w", err)
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
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("delete daily side cash: %w", err)
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