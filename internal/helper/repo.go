package helper

import "gorm.io/gorm"

type HelperRepo interface {
    Create(helper *Helper) error
    ExistsByMobile(mobile string) (bool, error)
    Get(id uint) (*Helper, error)
    Update(helper *Helper) error
    Delete(id uint) error
    GetAll(offset, limit int) ([]Helper, int64, error)
}

type helperRepo struct {
    db *gorm.DB 
}

func NewHelperRepo(db *gorm.DB) HelperRepo {
    return &helperRepo{db: db} 
}

func (r *helperRepo) Create(helper *Helper) error {
    return r.db.Create(helper).Error 
}

func (r *helperRepo) ExistsByMobile(mobile string) (bool, error) {
    var exists bool 
    query := "SELECT EXISTS(SELECT 1 FROM helpers WHERE mobile = ?)"
    if err := r.db.Raw(query, mobile).Scan(&exists).Error; err != nil {
        return false, err 
    }
    return exists, nil 
}

func (r *helperRepo) Get(id uint) (*Helper, error) {
    var helper Helper
    if err := r.db.First(&helper, id).Error; err != nil {
        return nil, err
    }
    return &helper, nil
}

func (r *helperRepo) Update(helper *Helper) error {
    return r.db.Save(helper).Error
}

func (r *helperRepo) Delete(id uint) error {
    return r.db.Delete(&Helper{}, id).Error
}

func (r *helperRepo) GetAll(offset, limit int) ([]Helper, int64, error) {
    var helpers []Helper
    var total int64

    if err := r.db.Model(&Helper{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&helpers).Error; err != nil {
        return nil, 0, err
    }

    return helpers, total, nil
}