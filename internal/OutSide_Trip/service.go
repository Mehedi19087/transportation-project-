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
    res.LoadPoint = req.LoadPoint
    res.UnloadPoint = req.UnloadPoint
    res.Rent = req.Rent
    res.Advance = req.Advance
    res.TripCost = req.TripCost
    res.Diesel = req.Diesel
    res.ExtraCost = req.ExtraCost
    res.DieselTaka = req.DieselTaka
    res.Pamp = req.Pamp
    res.Commission = req.Commission
    res.Month = req.Month
    res.VehicleName = req.VehicleName
    res.VehicleNumber = req.VehicleNumber
    res.DriverName = req.DriverName
    res.DriverPhone = req.DriverPhone
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