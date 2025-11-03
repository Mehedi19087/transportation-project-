package customer

import (
	"errors"
	"time"  
	"fmt"
)

type CustomerService interface {
	 CreateCustomer(req *CreateCustomerRequest) error  
	 UpdateCustomer(req *UpdateCustomerRequest, id uint ) error 
	 GetCustomer(id uint) (*Customer, error)
	 ListCustomers(page, pageSize int) ([]Customer, int64, error)
	 DeleteCustomer(id uint) error 
}

type customerService struct {
	 repo CustomerRepo
}

func NewCustomerService (repo CustomerRepo) CustomerService {
	 return &customerService {repo : repo}
}

func (s *customerService) CreateCustomer(req *CreateCustomerRequest) error {
	 if req.CustomerName == "" {
		 return errors.New("customer name is required")
	 }
	 customer:= &Customer {
		 CustomerName: req.CustomerName,
		 Mobile: req.Mobile,
		 Email:          req.Email,
         Address:        req.Address,
         OpeningBalance: req.OpeningBalance,
         Status:         req.Status,
		 CreatedAt: time.Now().UTC(),
	 }
	 err := s.repo.Create(customer)
	 if err != nil {
		 return fmt.Errorf("create customer: %w", err)
	 }
	 return nil 
}

func (s *customerService) UpdateCustomer(req *UpdateCustomerRequest, id uint) error {
	if id == 0 {
		return errors.New("invalid customer id")
	}

	customer, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	if req.CustomerName != "" {
		customer.CustomerName = req.CustomerName
	}
	customer.Mobile = req.Mobile
	customer.Email = req.Email
	customer.Address = req.Address
	customer.OpeningBalance = req.OpeningBalance
	customer.Status = req.Status
	customer.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(customer); err != nil {
		return err
	}
	return nil
}


func (s *customerService) GetCustomer(id uint) (*Customer, error) {
	  if id == 0 {
		  return nil , errors.New("invalid customer id")
	  }
	  customer , err := s.repo.Get(id) 
	  if err != nil {
		  return nil, err 
	  }
	  return customer, nil 
}

func(s *customerService) ListCustomers(page, pageSize int) ([]Customer, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}
	offset := (page - 1) * pageSize

	customers, total, err := s.repo.List(offset, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("list customers: %w", err)
	}

	return customers, total, nil
}

func(s *customerService) DeleteCustomer(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		 return err 
	}
	return nil 
}