package customer

import "gorm.io/gorm"

type CustomerRepo interface {
	 Create(customer *Customer) error 
	 Update(Customer *Customer) error 
	 Get(id uint) (*Customer, error)
	 Delete(id uint) error 
	 List(offset, limit int) ([]Customer, int64, error)
}

type customerRepo struct {
	 db *gorm.DB 
}

func NewCustomerRepo(db *gorm.DB) CustomerRepo {
	 return &customerRepo{db : db}
}

func (r *customerRepo) Create(customer *Customer) error {
	 return r.db.Create(customer).Error
}

func (r *customerRepo) Update(customer *Customer) error {
	 return r.db.Save(customer).Error 
}

func (r *customerRepo) Get(id uint)(*Customer, error){
	 var customer Customer 
	 err := r.db.First(&customer, id).Error 
	 if err != nil {
		 return nil , err 
	 }
	 return &customer, nil 
}

func (r *customerRepo) Delete(id uint) error {
	 return r.db.Delete(&Customer{}, id).Error 
}

func (r *customerRepo) List(offset, limit int) ([]Customer, int64, error) {
	var customers []Customer
	var total int64

	if err := r.db.Model(&Customer{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}