package vehicle

import (
    "errors"
    "gorm.io/gorm"
)

type VehicleRepo interface {
    Create(vehicle *Vehicle) error
    Update(vehicle *Vehicle) error
    Get(id uint) (*Vehicle, error)
    GetAll(offset, limit int) ([]Vehicle, int64, error)
    Delete(id uint) error
}

type vehicleRepo struct {
    db *gorm.DB
}

func NewVehicleRepo(db *gorm.DB) VehicleRepo {
    return &vehicleRepo{db: db}
}

func (r *vehicleRepo) Create(vehicle *Vehicle) error {
    return r.db.Create(vehicle).Error
}

func (r *vehicleRepo) Update(vehicle *Vehicle) error {
    return r.db.Save(vehicle).Error
}

func (r *vehicleRepo) Get(id uint) (*Vehicle, error) {
    var res Vehicle
    err := r.db.First(&res, id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("vehicle not found")
        }
        return nil, err
    }
    return &res, nil
}

func (r *vehicleRepo) Delete(id uint) error {
    return r.db.Delete(&Vehicle{}, id).Error
}

func (r *vehicleRepo) GetAll(offset, limit int) ([]Vehicle, int64, error) {
    var vehicles []Vehicle
    var total int64
    if err := r.db.Model(&Vehicle{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&vehicles).Error; err != nil {
        return nil, 0, err
    }
    return vehicles, total, nil
}