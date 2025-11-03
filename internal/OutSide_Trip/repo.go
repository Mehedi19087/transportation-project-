package outsidetrip

import (
    "errors"

    "gorm.io/gorm"
)

type OutSideTripRepo interface {
    Create(trip *OutSideTrip) error
    Update(trip *OutSideTrip) error
    GetAll(offset, limit int) ([]OutSideTrip, int64, error)
    Get(id uint) (*OutSideTrip, error)
    Delete(id uint) error
	GetByVehicleMonth(vehicleNumber, month string) ([]OutSideTrip, int64, error)
}

type outSideTripRepo struct {
    db *gorm.DB
}

func NewOutSideTripRepo(db *gorm.DB) OutSideTripRepo {
    return &outSideTripRepo{db: db}
}

func (r *outSideTripRepo) Create(trip *OutSideTrip) error {
    return r.db.Create(trip).Error
}

func (r *outSideTripRepo) Update(trip *OutSideTrip) error {
    return r.db.Save(trip).Error
}

func (r *outSideTripRepo) Get(id uint) (*OutSideTrip, error) {
    var res OutSideTrip
    err := r.db.First(&res, id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("outside trip not found")
        }
        return nil, err
    }
    return &res, nil
}

func (r *outSideTripRepo) Delete(id uint) error {
    return r.db.Delete(&OutSideTrip{}, id).Error
}

func (r *outSideTripRepo) GetAll(offset, limit int) ([]OutSideTrip, int64, error) {
    var trips []OutSideTrip
    var total int64
    if err := r.db.Model(&OutSideTrip{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&trips).Error; err != nil {
        return nil, 0, err
    }
    return trips, total, nil
}

func (r *outSideTripRepo) GetByVehicleMonth(vehicleNumber, month string) ([]OutSideTrip, int64, error) {
    var trips []OutSideTrip
    var total int64
    
    if err := r.db.Model(&OutSideTrip{}).Where("vehicle_number = ? AND month = ?", vehicleNumber, month).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    
    if err := r.db.Where("vehicle_number = ? AND month = ?", vehicleNumber, month).Find(&trips).Error; err != nil {
        return nil, 0, err
    }
    
    return trips, total, nil
}