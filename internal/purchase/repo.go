package purchase

import (
    "errors"

    "gorm.io/gorm"
)

type PurchaseRepo interface {
    Create(p *Purchase) error
    Update(p *Purchase) error
    Get(id uint) (*Purchase, error)
    GetAll(offset, limit int) ([]Purchase, int64, error)
    Delete(id uint) error
}

type purchaseRepo struct {
    db *gorm.DB
}

func NewPurchaseRepo(db *gorm.DB) PurchaseRepo {
    return &purchaseRepo{db: db}
}

func (r *purchaseRepo) Create(p *Purchase) error {
    return r.db.Create(p).Error
}

func (r *purchaseRepo) Update(p *Purchase) error {
    return r.db.Save(p).Error
}

func (r *purchaseRepo) Get(id uint) (*Purchase, error) {
    var res Purchase
    err := r.db.First(&res, id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("purchase not found")
        }
        return nil, err
    }
    return &res, nil
}

func (r *purchaseRepo) Delete(id uint) error {
    return r.db.Delete(&Purchase{}, id).Error
}

func (r *purchaseRepo) GetAll(offset, limit int) ([]Purchase, int64, error) {
    var items []Purchase
    var total int64
    if err := r.db.Model(&Purchase{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&items).Error; err != nil {
        return nil, 0, err
    }
    return items, total, nil
}