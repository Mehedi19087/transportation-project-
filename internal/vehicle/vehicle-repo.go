package vehicle

import (
    "errors"
    "transportation/internal/utils"

    "gorm.io/gorm"
)

type VehicleRepo interface {
    Create(vehicle *Vehicle) error
    Update(vehicle *Vehicle) error
    Get(id uint) (*Vehicle, error)
    GetAll(cursor *utils.Cursor, limit int) ([]Vehicle, error)
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

func (r *vehicleRepo) GetAll(cursor *utils.Cursor, limit int) ([]Vehicle, error) {
    var vehicles []Vehicle
    
    query := r.db.Model(&Vehicle{})

    if cursor != nil {
        query = query.Where("created_at < ? OR (created_at = ? AND id < ?)", cursor.LastCreatedAt, cursor.LastCreatedAt, cursor.LastID)
    }

    if err := query.Order("created_at DESC, id DESC").Limit(limit + 1).Find(&vehicles).Error; err != nil {
        return nil, err
    }
    return vehicles, nil
}