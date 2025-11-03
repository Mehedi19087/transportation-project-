package driver

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)


type DriverService interface {
	 CreateDriver(req *CreateDriver) error
	 GetDriver(id uint) (*Driver, error)
     UpdateDriver(id uint, req *UpdateDriver) error
     DeleteDriver(id uint) error
     GetAllDriver(page, pageSize int) ([]Driver, int64, error)
}

type driverService struct {
	 repo DriverRepo
}

func NewDriverService(repo DriverRepo) DriverService {
	 return &driverService{repo : repo}
}

func(s *driverService) CreateDriver(req *CreateDriver) error {
	  if req.DriverName == ""{
		 return errors.New("driver name is required")
	  }
	  fmt.Printf("Received NID: '%s'\n", req.NID)
	  log.Println(req.NID)
	  driver:= &Driver {
		 Name: req.DriverName,
		 Mobile: req.Mobile,
		 Address: req.Address,
		 Emergency: req.Emergency,
		 License: req.License,
		 Expired: req.Expired,
		 NID: req.NID,
		 JoiningDate: req.JoiningDate,
		 ImageUrl: req.ImageURL,
	  }
	  err:= s.repo.Create(driver)
	  if err != nil {
		 return err 
	  }
	  return nil 
}

func (s *driverService) GetDriver(id uint) (*Driver, error) {
    res, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *driverService) UpdateDriver(id uint, req *UpdateDriver) error {
    res, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("driver data is missing")
        }
        return err
    }
    
    if req.DriverName != "" {
        res.Name = req.DriverName
    }
    if req.Mobile != "" {
        res.Mobile = req.Mobile
    }
    if req.Address != "" {
        res.Address = req.Address
    }
    if req.Emergency != "" {
        res.Emergency = req.Emergency
    }
    if req.License != "" {
        res.License = req.License
    }
    if !req.Expired.IsZero() {
        res.Expired = req.Expired
    }
    if req.NID != "" {
        res.NID = req.NID
    }
    if !req.JoiningDate.IsZero() {
        res.JoiningDate = req.JoiningDate
    }
    if req.Status != "" {
        res.Status = req.Status
    }
	if req.ImageURL != nil {
        res.ImageUrl = req.ImageURL
    }
    
    if err := s.repo.Update(res); err != nil {
        return err
    }
    return nil
}

func (s *driverService) DeleteDriver(id uint) error {
    err := s.repo.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *driverService) GetAllDriver(page, pageSize int) ([]Driver, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize

    drivers, total, err := s.repo.GetAll(offset, pageSize)

    if err != nil {
        return nil, 0, fmt.Errorf("list drivers: %w", err)
    }
    return drivers, total, nil
}