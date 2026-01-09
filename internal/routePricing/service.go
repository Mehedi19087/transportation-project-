package routepricing

import (
    "errors"
    "fmt"
    "gorm.io/gorm"
)

type RoutePricingService interface {
    CreateRoutePricing(req *RoutePricingReq) error
    GetRoutePricing(id uint) (*RoutePricing, error)
    UpdateRoutePricing(id uint, req *RoutePricingUpdateReq) error
    DeleteRoutePricing(id uint) error
    GetAllRoutePricing(page, pageSize int) ([]RoutePricing, int64, error)
    GetRate(dealerName, destination string) (float64, error)
    GetDealerNames() ([]string, error)
}

type routePricingService struct {
    repo RoutePricingRepo
}

func NewRoutePricingService(repo RoutePricingRepo) RoutePricingService {
    return &routePricingService{repo: repo}
}

func (s *routePricingService) CreateRoutePricing(req *RoutePricingReq) error {
    if req.DealerName == "" {
        return errors.New("dealer name is required")
    }
    if req.Destination == "" {
        return errors.New("destination is required")
    }

    routePricing := &RoutePricing{
        DealerName:  req.DealerName,
        Destination: req.Destination,
        Rate:        req.Rate,
    }

    if err := s.repo.Create(routePricing); err != nil {
        return fmt.Errorf("failed to create route pricing: %w", err)
    }

    return nil
}

func (s *routePricingService) GetRoutePricing(id uint) (*RoutePricing, error) {
    res, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *routePricingService) UpdateRoutePricing(id uint, req *RoutePricingUpdateReq) error {
    res, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("route pricing data is missing")
        }
        return err
    }

    if req.DealerName != "" {
        res.DealerName = req.DealerName
    }
    if req.Destination != "" {
        res.Destination = req.Destination
    }
    if req.Rate != 0 {
        res.Rate = req.Rate
    }

    if err := s.repo.Update(res); err != nil {
        return err
    }
    return nil
}

func (s *routePricingService) DeleteRoutePricing(id uint) error {
    err := s.repo.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *routePricingService) GetAllRoutePricing(page, pageSize int) ([]RoutePricing, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize

    routePricings, total, err := s.repo.GetAll(offset, pageSize)

    if err != nil {
        return nil, 0, fmt.Errorf("list route pricings: %w", err)
    }
    return routePricings, total, nil
}

func (s *routePricingService) GetRate(dealerName, destination string) (float64, error) {
    rate, err := s.repo.GetRate(dealerName, destination)
    if err != nil {
        return 0, errors.New("rate not found for the given dealer and destination")
    }
    return rate, nil
}

func (s *routePricingService) GetDealerNames() ([]string, error) {
    return s.repo.GetDealerNames()
}