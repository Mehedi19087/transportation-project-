package calan

import (
	"gorm.io/gorm"
)

type CalanService interface {
	CreateCalan(tx *gorm.DB, calan *Calan) error
	GetCalans(filters FilterCalanReq) ([]Calan, int64, error)
	UpdateCalanStatus(id uint, status string) error
}

type calanService struct {
	repo CalanRepo
}

func NewCalanService(repo CalanRepo) CalanService {
	return &calanService{repo: repo}
}

func (s *calanService) CreateCalan(tx *gorm.DB, calan *Calan) error {
	// Business logic validation can go here if needed
	return s.repo.Create(tx, calan)
}

func (s *calanService) GetCalans(filters FilterCalanReq) ([]Calan, int64, error) {
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.PageSize < 1 || filters.PageSize > 100 {
		filters.PageSize = 10
	}
	return s.repo.GetAll(filters)
}

func (s *calanService) UpdateCalanStatus(id uint, status string) error {
	return s.repo.UpdateStatus(id, status)
}
