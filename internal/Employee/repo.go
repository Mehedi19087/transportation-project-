package employee

import "gorm.io/gorm"

type EmployeeRepo interface {
	Create(employee *Employee) error
	Update(employee *Employee) error 
	Get(id uint) (*Employee, error)
	Delete(id uint) error 
	GetAll(limit, offset int) ([]Employee, int64, error)
}

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) EmployeeRepo {
	 return &employeeRepo{db : db}
}

func(r *employeeRepo) Create(employee *Employee) error {
	 return r.db.Create(employee).Error
}

func(r *employeeRepo) Update(employee *Employee) error {
	 return r.db.Save(employee).Error
}

func(r *employeeRepo) Get(id uint) (*Employee, error) {
	 var employee Employee 
	 if err:= r.db.First(&employee, id).Error; err != nil {
		 return nil, err
	 }
	 return &employee, nil 
}

func(r *employeeRepo) Delete(id uint) error {
	 return r.db.Delete(&Employee{}, id).Error
}

func(r *employeeRepo) GetAll(limit, offset int) ([]Employee, int64, error) {
	 var total int64 
	 var employees []Employee

	 if err := r.db.Model(&Employee{}).Count(&total).Error; err!= nil {
		return nil, 0, err 
	 }

	 if err := r.db.Order("created_at DESC").Offset(offset).Limit(limit).Find(&employees).Error; err != nil {
		 return nil, 0, err 
	 }
	 return employees, total, nil 
}