package trip

import (
    "gorm.io/gorm"
)

type Repository interface {
    Create(t *Trip) error
    Get(id uint) (*Trip, error)
    GetByProduct(productID uint, offset, limit int) ([]Trip, int64, error)
    Update(t *Trip) error
    Delete(id uint) error
}

type repo struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repo{db: db}
}

func (r *repo) Create(t *Trip) error {
    return r.db.Create(t).Error
}

func (r *repo) Get(id uint) (*Trip, error) {
    var res Trip
    if err := r.db.First(&res, id).Error; err != nil {
        return nil, err
    }
    return &res, nil
}

func (r *repo) GetByProduct(productID uint, offset, limit int) ([]Trip, int64, error) {
    var total int64
    var items []Trip
    if err := r.db.Model(&Trip{}).Where("product_id = ?", productID).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    if err := r.db.Where("product_id = ?", productID).
        Order("created_at DESC").
        Limit(limit).Offset(offset).
        Find(&items).Error; err != nil {
        return nil, 0, err
    }
    return items, total, nil
}

func (r *repo) Update(t *Trip) error {
    return r.db.Save(t).Error
}

func (r *repo) Delete(id uint) error {
    return r.db.Delete(&Trip{}, id).Error
}