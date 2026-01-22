package calan

import (
	"gorm.io/gorm"
)

type CalanRepo interface {
	Create(tx *gorm.DB, calan *Calan) error
	GetAll(filters FilterCalanReq) ([]Calan, int64, error)
	UpdateStatus(id uint, status string) error
    GetByID(id uint) (*Calan, error)
	DeleteByTripID(tx *gorm.DB, tripID uint) error
	DeleteByOutsideTripID(tx *gorm.DB, outsideTripID uint) error
}

type calanRepo struct {
	db *gorm.DB
}

func NewCalanRepo(db *gorm.DB) CalanRepo {
	return &calanRepo{db: db}
}

// Create accepts a tx (transaction) to allow atomic operations with Trip/OutsideTrip
func (r *calanRepo) Create(tx *gorm.DB, calan *Calan) error {
	var db *gorm.DB
	if tx != nil {
		db = tx
	} else {
		db = r.db
	}
	return db.Create(calan).Error
}

func (r *calanRepo) GetAll(filters FilterCalanReq) ([]Calan, int64, error) {
	var calans []Calan
	var total int64

	query := r.db.Model(&Calan{})

	if filters.Date != "" {
		query = query.Where("date = ?", filters.Date)
	}
	if filters.VehicleNo != "" {
		query = query.Where("vehicle_no ILIKE ?", "%"+filters.VehicleNo+"%")
	}
	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.PageSize
	err := query.Order("created_at DESC").
		Limit(filters.PageSize).
		Offset(offset).
		Find(&calans).Error

	return calans, total, err
}

func (r *calanRepo) UpdateStatus(id uint, status string) error {
	return r.db.Model(&Calan{}).Where("id = ?", id).Update("status", status).Error
}

func (r *calanRepo) GetByID(id uint) (*Calan, error) {
    var calan Calan
    err := r.db.First(&calan, id).Error
    if err != nil {
        return nil, err
    }
    return &calan, nil
}

func (r *calanRepo) DeleteByTripID(tx *gorm.DB, tripID uint) error {
	return tx.Where("trip_id = ?", tripID).Delete(&Calan{}).Error
}

func (r *calanRepo) DeleteByOutsideTripID(tx *gorm.DB, outsideTripID uint) error {
	return tx.Where("outside_trip_id = ?", outsideTripID).Delete(&Calan{}).Error
}
