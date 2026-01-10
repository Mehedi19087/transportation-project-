package outsidetrip

import (
    "errors"
    "fmt"
    "time"

    "gorm.io/gorm"
)

type OutSideTripService interface {
    CreateOutSideTrip(req *OutSideTripReq) error
    GetOutSideTrip(id uint) (*OutSideTrip, error)
    UpdateOutSideTrip(id uint, req *OutSideTripUpdateReq) error
    DeleteOutSideTrip(id uint) error
    GetAllOutSideTrip(page, pageSize int) ([]OutSideTrip, int64, error)
	GetVehicleMonthlySummary(vehicleNumber, month string) (*VehicleMonthlySummary, error)
}

type outSideTripService struct {
    repo OutSideTripRepo
}

func NewOutSideTripService(repo OutSideTripRepo) OutSideTripService {
    return &outSideTripService{repo: repo}
}

func (s *outSideTripService) CreateOutSideTrip(req *OutSideTripReq) error {
    if req.LoadPoint == "" {
        return errors.New("load point is required")
    }
    trip := &OutSideTrip{
        Date:          req.Date,
        LoadPoint:     req.LoadPoint,
        UnloadPoint:   req.UnloadPoint,
        Rent:          req.Rent,
        Advance:       req.Advance,
        TripCost:      req.TripCost,
        Diesel:        req.Diesel,
        ExtraCost:     req.ExtraCost,
        DieselTaka:    req.DieselTaka,
        Pamp:          req.Pamp,
        Commission:    req.Commission,
        Month:         req.Month,
        VehicleName:   req.VehicleName,
        VehicleNumber: req.VehicleNumber,
        DriverName:    req.DriverName,
        DriverPhone:   req.DriverPhone,
        Due:           req.Due,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }

    if err := s.repo.Create(trip); err != nil {
        return fmt.Errorf("failed to create outside trip: %w", err)
    }

    return nil
}

func (s *outSideTripService) GetOutSideTrip(id uint) (*OutSideTrip, error) {
    res, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *outSideTripService) UpdateOutSideTrip(id uint, req *OutSideTripUpdateReq) error {
    res, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("outside trip data is missing")
        }
        return err
    }
    if req.Date != nil {
        res.Date = *req.Date
    }
    if req.LoadPoint != nil {
        res.LoadPoint = *req.LoadPoint
    }
    if req.UnloadPoint != nil {
        res.UnloadPoint = *req.UnloadPoint
    }
    if req.Rent != nil {
        res.Rent = *req.Rent
    }
    if req.Advance != nil {
        res.Advance = *req.Advance
    }
    if req.TripCost != nil {
        res.TripCost = *req.TripCost
    }
    if req.Diesel != nil {
        res.Diesel = *req.Diesel
    }
    if req.ExtraCost != nil {
        res.ExtraCost = *req.ExtraCost
    }
    if req.DieselTaka != nil {
        res.DieselTaka = *req.DieselTaka
    }
    if req.Pamp != nil {
        res.Pamp = *req.Pamp
    }
    if req.Commission != nil {
        res.Commission = *req.Commission
    }
    if req.Month != nil {
        res.Month = *req.Month
    }
    if req.VehicleName != nil {
        res.VehicleName = *req.VehicleName
    }
    if req.VehicleNumber != nil {
        res.VehicleNumber = *req.VehicleNumber
    }
    if req.DriverName != nil {
        res.DriverName = *req.DriverName
    }
    if req.DriverPhone != nil {
        res.DriverPhone = *req.DriverPhone
    }
    if req.Due != nil {
        res.Due = *req.Due
    }
    if req.DueStatus != nil {
        res.DueStatus = *req.DueStatus
    }
    
    res.UpdatedAt = time.Now()

    if err := s.repo.Update(res); err != nil {
        return err
    }
    return nil
}

func (s *outSideTripService) DeleteOutSideTrip(id uint) error {
    err := s.repo.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *outSideTripService) GetAllOutSideTrip(page, pageSize int) ([]OutSideTrip, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 10 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize

    trips, total, err := s.repo.GetAll(offset, pageSize)

    if err != nil {
        return nil, 0, fmt.Errorf("list outside trips: %w", err)
    }
    return trips, total, nil
}

func(s *outSideTripService) GetVehicleMonthlySummary(vehicleNumber, month string) (*VehicleMonthlySummary, error) {
	if vehicleNumber == "" {
        return nil, errors.New("vehicle number is required")
    }
    if month == "" {
        return nil, errors.New("month is required")
    }
    
    trips, total, err := s.repo.GetByVehicleMonth(vehicleNumber, month)
    if err != nil {
        return nil, fmt.Errorf("failed to get vehicle monthly trips: %w", err)
    }

	summary := &VehicleMonthlySummary{
        VehicleNumber: vehicleNumber,
        Month:         month,
        TripCount:     total,
    }

	for _, trip := range trips {
        summary.TotalRent += trip.Rent
        summary.TotalAdvance += trip.Advance
        summary.TotalTripCost += trip.TripCost
        summary.TotalDiesel += trip.Diesel
        summary.TotalExtraCost += trip.ExtraCost
        summary.TotalDieselTaka += trip.DieselTaka
        summary.TotalCommission += trip.Commission
    }
	return summary, nil 
}