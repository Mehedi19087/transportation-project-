package trip

import (
    "errors"
    "time"

    "gorm.io/gorm"
)

type Service interface {
    CreateTrip(req *CreateTripReq) error
    GetTrip(id uint) (*Trip, error)
    GetProductTrips(productID uint, page, pageSize int) ([]Trip, int64, error)
    UpdateTrip(id uint, req *UpdateTripReq) error
    DeleteTrip(id uint) error
}

type service struct {
    repo Repository
}

func NewService(r Repository) Service {
    return &service{repo: r}
}

func (s *service) CreateTrip(req *CreateTripReq) error {
    if req.ProductID == 0 {
        return errors.New("product_id is required")
    }
    item := &Trip{
        ProductID:     req.ProductID,
        BrandName:     req.BrandName,
        Category:      req.Category,
        Date:          req.Date,
        TripType:      req.TripType,
        TripNo:        req.TripNo,
        InvoiceNo:     req.InvoiceNo,
        VehicleName:   req.VehicleName,
        VehicleNo:     req.VehicleNo,
        EngineNo:      req.EngineNo,
        ChassisNo:     req.ChassisNo,
        DriverName:    req.DriverName,
        DriverMobile:  req.DriverMobile,
        HelperName:    req.HelperName,
        LoadPoint:     req.LoadPoint,
        UnloadPoint:   req.UnloadPoint,
        Destination:   req.Destination,
        Route:         req.Route,
        District:      req.District,
        Quantity:      req.Quantity,
        UnitPrice:     req.UnitPrice,
        TotalRate:     req.TotalRate,
        Cash:          req.Cash,
        Advance:       req.Advance,
        Due:           req.Due,
        BillNo:        req.BillNo,
        BillDate:      req.BillDate,
        PaymentType:   req.PaymentType,
        TruckSize:     req.TruckSize,
        Weight:        req.Weight,
        FuelType:      req.FuelType,
        FuelCost:      req.FuelCost,
        TransportType: req.TransportType,
        Remarks:       req.Remarks,
        Status:        req.Status,
        CreatedBy:     req.CreatedBy,
        ApprovedBy:    req.ApprovedBy,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }
    return s.repo.Create(item)
}

func (s *service) GetTrip(id uint) (*Trip, error) {
    return s.repo.Get(id)
}

func (s *service) GetProductTrips(productID uint, page, pageSize int) ([]Trip, int64, error) {
    if productID == 0 {
        return nil, 0, errors.New("invalid product id")
    }
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize
    return s.repo.GetByProduct(productID, offset, pageSize)
}

func (s *service) UpdateTrip(id uint, req *UpdateTripReq) error {
    item, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("trip not found")
        }
        return err
    }

    if req.ProductID != nil {
        item.ProductID = *req.ProductID
    }
    if req.BrandName != nil {
        item.BrandName = req.BrandName
    }
    if req.Category != nil {
        item.Category = req.Category
    }
    if req.Date != nil {
        item.Date = req.Date
    }
    if req.TripType != nil {
        item.TripType = req.TripType
    }
    if req.TripNo != nil {
        item.TripNo = req.TripNo
    }
    if req.InvoiceNo != nil {
        item.InvoiceNo = req.InvoiceNo
    }
    if req.VehicleName != nil {
        item.VehicleName = req.VehicleName
    }
    if req.VehicleNo != nil {
        item.VehicleNo = req.VehicleNo
    }
    if req.EngineNo != nil {
        item.EngineNo = req.EngineNo
    }
    if req.ChassisNo != nil {
        item.ChassisNo = req.ChassisNo
    }
    if req.DriverName != nil {
        item.DriverName = req.DriverName
    }
    if req.DriverMobile != nil {
        item.DriverMobile = req.DriverMobile
    }
    if req.HelperName != nil {
        item.HelperName = req.HelperName
    }
    if req.LoadPoint != nil {
        item.LoadPoint = req.LoadPoint
    }
    if req.UnloadPoint != nil {
        item.UnloadPoint = req.UnloadPoint
    }
    if req.Destination != nil {
        item.Destination = req.Destination
    }
    if req.Route != nil {
        item.Route = req.Route
    }
    if req.District != nil {
        item.District = req.District
    }
    if req.Quantity != nil {
        item.Quantity = req.Quantity
    }
    if req.UnitPrice != nil {
        item.UnitPrice = req.UnitPrice
    }
    if req.TotalRate != nil {
        item.TotalRate = req.TotalRate
    }
    if req.Cash != nil {
        item.Cash = req.Cash
    }
    if req.Advance != nil {
        item.Advance = req.Advance
    }
    if req.Due != nil {
        item.Due = req.Due
    }
    if req.BillNo != nil {
        item.BillNo = req.BillNo
    }
    if req.BillDate != nil {
        item.BillDate = req.BillDate
    }
    if req.PaymentType != nil {
        item.PaymentType = req.PaymentType
    }
    if req.TruckSize != nil {
        item.TruckSize = req.TruckSize
    }
    if req.Weight != nil {
        item.Weight = req.Weight
    }
    if req.FuelType != nil {
        item.FuelType = req.FuelType
    }
    if req.FuelCost != nil {
        item.FuelCost = req.FuelCost
    }
    if req.TransportType != nil {
        item.TransportType = req.TransportType
    }
    if req.Remarks != nil {
        item.Remarks = req.Remarks
    }
    if req.Status != nil {
        item.Status = req.Status
    }
    if req.CreatedBy != nil {
        item.CreatedBy = req.CreatedBy
    }
    if req.ApprovedBy != nil {
        item.ApprovedBy = req.ApprovedBy
    }
    item.UpdatedAt = time.Now()

    return s.repo.Update(item)
}

func (s *service) DeleteTrip(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("trip not found")
        }
        return err
    }
    return s.repo.Delete(id)
}