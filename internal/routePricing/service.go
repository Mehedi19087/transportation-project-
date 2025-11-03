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
}

type routePricingService struct {
    repo RoutePricingRepo
}

func NewRoutePricingService(repo RoutePricingRepo) RoutePricingService {
    return &routePricingService{repo: repo}
}

func (s *routePricingService) CreateRoutePricing(req *RoutePricingReq) error {
    if req.CustomerName == "" {
        return errors.New("customer name is required")
    }
    if req.LoadPoint == "" {
        return errors.New("load point is required")
    }
    if req.UnloadPoint == "" {
        return errors.New("unload point is required")
    }

    routePricing := &RoutePricing{
        CustomerName: req.CustomerName,
        Size:         req.Size,
        LoadPoint:    req.LoadPoint,
        UnloadPoint:  req.UnloadPoint,
        Rate:         req.Rate,
        Weight:       req.Weight,
        Notes:        req.Notes,
        Status:       req.Status,
        VehicleCategory: req.VehicleCategory,
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

    res.CustomerName = req.CustomerName
    res.Size = req.Size
    res.LoadPoint = req.LoadPoint
    res.UnloadPoint = req.UnloadPoint
    res.Rate = req.Rate
    res.Weight = req.Weight
    res.Notes = req.Notes
    res.Status = req.Status
    res.VehicleCategory= req.VehicleCategory

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