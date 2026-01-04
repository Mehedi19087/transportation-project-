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

	record := &DailySideCash{
		Date:             d,
		ProductID:        productID,
		Cash:             req.Cash,
		RemainingBalance: req.RemainingBalance,
		TripCost:         req.TripCost,
		OtherCost:        req.OtherCost,
		OtherCostDetails: req.OtherCostDetails,
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
		existing.Cash = *req.Cash
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