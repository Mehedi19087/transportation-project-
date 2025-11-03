package dealer

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type DealerService interface {
	CreateDealer(req *DealerReq) error
	GetDealer(id uint) (*Dealer, error)
	UpdateDealer(id uint , req *DealerUpdateReq) error 
	DeleteDealer(id uint) error
	GetAllDealer(page , pageSize int) ([]Dealer, int64, error)
}

type dealerService struct {
	repo DealerRepo
}

func NewDealerService(repo DealerRepo) DealerService {
	return &dealerService{repo: repo}
}

func (s *dealerService) CreateDealer(req *DealerReq) error {
	if req.DealerName == "" {
		return errors.New("dealer name is required")
	}
	dealer := &Dealer{
		Name:        req.DealerName,
		Destination: req.Destination,
		Status:      req.Status,
	}
	
	// Save to database
	if err := s.repo.Create(dealer); err != nil {
		return fmt.Errorf("failed to create dealer: %w", err)
	}
	
	return nil
}
func (s *dealerService) GetDealer(id uint) (*Dealer, error) {
	res , err := s.repo.Get(id)
	if err != nil {
		 return nil, err 
	}
	return res , nil 
}

func (s *dealerService) UpdateDealer(id uint , req *DealerUpdateReq) error {
	res, err := s.repo.Get(id) 
	if err != nil {
		 if err == gorm.ErrRecordNotFound {
			 return errors.New("dealer data is missing")
		 }
		 return err 
	}
	res.Name= req.DealerName 
	res.Destination=req.Destination
	res.Status=req.Status
	if err:=s.repo.Update(res) ; err != nil {
		 return  err 
	}
    return nil 
}

func(s *dealerService) DeleteDealer(id uint) error {
	 err:= s.repo.Delete(id) 
	 if err!= nil {
		 return err 
	 }
	 return nil 
}

func(s *dealerService) GetAllDealer(page , pageSize int) ([]Dealer, int64, error) {
	 if page < 1 {
		 page =1 
	 }
	 if pageSize < 1 || pageSize>10 {
		 pageSize=10 
	 }
	 offset:= (page-1)*pageSize

	 dealers, total, err := s.repo.GetAll(offset, pageSize)

	 if err != nil {
        return nil, 0, fmt.Errorf("list customers: %w", err)
    }
    return dealers, total, nil
}