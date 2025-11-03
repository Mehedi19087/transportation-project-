package routepricing

import (
    "errors"
    "gorm.io/gorm"
)

type RoutePricingRepo interface {
    Create(routePricing *RoutePricing) error
    Update(routePricing *RoutePricing) error
    GetAll(offset, limit int) ([]RoutePricing, int64, error)
    Get(id uint) (*RoutePricing, error)
    Delete(id uint) error
}

type routePricingRepo struct {
    db *gorm.DB
}

func NewRoutePricingRepo(db *gorm.DB) RoutePricingRepo {
    return &routePricingRepo{db: db}
}

func (r *routePricingRepo) Create(routePricing *RoutePricing) error {
    return r.db.Create(routePricing).Error
}

func (r *routePricingRepo) Update(routePricing *RoutePricing) error {
    return r.db.Save(routePricing).Error
}

func (r *routePricingRepo) Get(id uint) (*RoutePricing, error) {
    var res RoutePricing
    err := r.db.First(&res, id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("route pricing not found")
        }
        return nil, err
    }
    return &res, nil
}

func (r *routePricingRepo) Delete(id uint) error {
    return r.db.Delete(&RoutePricing{}, id).Error
}

func (r *routePricingRepo) GetAll(offset, limit int) ([]RoutePricing, int64, error) {
    var routePricings []RoutePricing
    var total int64
    if err := r.db.Model(&RoutePricing{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&routePricings).Error; err != nil {
        return nil, 0, err
    }
    return routePricings, total, nil
}