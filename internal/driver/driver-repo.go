package driver

import (
	"errors"

	"gorm.io/gorm"
)

type DriverRepo interface {
	 Create(driver *Driver) error 
	 Update(driver *Driver) error 
	 Get(id uint) (*Driver, error)
	 GetAll(offset , limit int) ([]Driver, int64, error) 
	 Delete(id uint) error 
}

type driverRepo struct {
	 db *gorm.DB 
}

func NewDriverRepo(db *gorm.DB) DriverRepo {
	 return &driverRepo{db : db}
}

func(r *driverRepo) Create(driver *Driver) error {
	 return r.db.Create(driver).Error
}

func(r *driverRepo) Update(driver *Driver) error {
	 return r.db.Save(driver).Error 
}

func(r *driverRepo) Get(id uint) (*Driver, error) {
	 var res Driver 

	 err := r.db.First(&res, id).Error 
	 if err != nil {
		  if err == gorm.ErrRecordNotFound {
			 return nil, errors.New("driver not found")
		  }
		  return nil , err 
	 }
	 return &res , nil 
}
func (r *driverRepo) Delete(id uint) error{
    return r.db.Delete(&Driver{}, id).Error
}

func (r *driverRepo) GetAll(offset, limit int) ([]Driver, int64, error) {
    var drivers []Driver
    var total int64
    if err := r.db.Model(&Driver{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }
    if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&drivers).Error; err != nil {
        return nil, 0, err
    }
    return drivers, total, nil
}