package ownvehicle

import (
    "errors"
    "fmt"
    "time"

    "gorm.io/gorm"
)

// ========== Legacy Service (keep as is) ==========

type Service interface {
    GetOwnVehiclesByDriverAndDate(driverName, startDate, endDate string) ([]OwnVehicle, error)
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo: repo}
}

func (s *service) GetOwnVehiclesByDriverAndDate(driverName, startDate, endDate string) ([]OwnVehicle, error) {
    loc, err := time.LoadLocation("Asia/Dhaka")
    if err != nil {
        loc = time.FixedZone("BDT", 6*60*60)
    }
    const layout = "2006-01-02"

    startLocal, err := time.ParseInLocation(layout, startDate, loc)
    if err != nil {
        return nil, fmt.Errorf("invalid start_date, expected YYYY-MM-DD: %w", err)
    }
    endLocal, err := time.ParseInLocation(layout, endDate, loc)
    if err != nil {
        return nil, fmt.Errorf("invalid end_date, expected YYYY-MM-DD: %w", err)
    }
    endLocal = endLocal.Add(24 * time.Hour).Add(-time.Nanosecond)

    startUTC := startLocal.In(time.UTC)
    endUTC := endLocal.In(time.UTC)
    trips, err := s.repo.GetTripsByDriverBetween(driverName, startUTC, endUTC)
    if err != nil {
        return nil, fmt.Errorf("fetch trips: %w", err)
    }

    result := make([]OwnVehicle, 0, len(trips))
    for _, t := range trips {
        lp := derefStr(t.LoadPoint)
        up := derefStr(t.UnloadPoint)
        adv := derefF64(t.Advance)

        rate, rateErr := s.repo.GetRateByLocations(lp, up)
        if rateErr != nil {
            rate = 0
        }

        ov := OwnVehicle{
            TripID:      t.ID,
            LoadPoint:   lp,
            UnloadPoint: up,
            Rent:        rate,
            Advance:     adv,
            Commission:  rate * 0.1,
        }
        _ = s.repo.CreateOwnVehicle(&ov)
        result = append(result, ov)
    }

    return result, nil
}

func derefStr(p *string) string {
    if p == nil {
        return ""
    }
    return *p
}

func derefF64(p *float64) float64 {
    if p == nil {
        return 0
    }
    return *p
}

// ========== OwnVehicleTrip Service ==========

type OwnVehicleTripService interface {
    Create(req *CreateOwnVehicleTripDTO) error
    Get(id uint) (*OwnVehicleTrip, error)
    GetAll(page, pageSize int) ([]OwnVehicleTrip, int64, error)
    Update(id uint, req *UpdateOwnVehicleTripDTO) error
    Delete(id uint) error
    GetByVehicleAndDateRange(vehicleNo, startDate, endDate string) ([]OwnVehicleTrip, error)
}

type ownVehicleTripService struct {
    repo OwnVehicleTripRepository
}

func NewOwnVehicleTripService(repo OwnVehicleTripRepository) OwnVehicleTripService {
    return &ownVehicleTripService{repo: repo}
}

func (s *ownVehicleTripService) Create(req *CreateOwnVehicleTripDTO) error {
    if req.VehicleNo == "" {
        return errors.New("vehicle number is required")
    }
    if req.Date == "" {
        return errors.New("date is required")
    }
    

    dieselTaka := req.Diesel * req.DieselPrice
    commission := req.Rent * 0.1
    profit := req.Rent - (req.TripCost + req.ExtraCost + dieselTaka + commission)

    ownVehicleTrip := &OwnVehicleTrip{
        Date:          req.Date,
        VehicleNo:     req.VehicleNo,
        LoadPoint:     req.LoadPoint,
        UnloadPoint:   req.UnloadPoint,
        Pump:          req.Pump,
        Rent:          req.Rent,
        Advance:       req.Advance,
        TripCost:      req.TripCost,
        Diesel:        req.Diesel,
        DieselPrice:   req.DieselPrice,
        DieselTaka:    dieselTaka,
        ExtraCost:     req.ExtraCost,
        Commission:    commission,
        Profit:        profit,
        TripID:        req.TripID,
        OutsideTripID: req.OutsideTripID,
    }

    if err := s.repo.Create(ownVehicleTrip); err != nil {
        return fmt.Errorf("failed to create own vehicle trip: %w", err)
    }

    return nil
}

func (s *ownVehicleTripService) Get(id uint) (*OwnVehicleTrip, error) {
    trip, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.New("own vehicle trip not found")
        }
        return nil, err
    }
    return trip, nil
}

func (s *ownVehicleTripService) GetAll(page, pageSize int) ([]OwnVehicleTrip, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    trips, total, err := s.repo.GetAll(offset, pageSize)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to get own vehicle trips: %w", err)
    }

    return trips, total, nil
}

func (s *ownVehicleTripService) Update(id uint, req *UpdateOwnVehicleTripDTO) error {
    trip, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("own vehicle trip not found")
        }
        return err
    }

    if req.Date != nil {
        trip.Date = *req.Date
    }
    if req.VehicleNo != nil {
        trip.VehicleNo = *req.VehicleNo
    }
    if req.LoadPoint != nil {
        trip.LoadPoint = *req.LoadPoint
    }
    if req.UnloadPoint != nil {
        trip.UnloadPoint = *req.UnloadPoint
    }
    if req.Pump != nil {
        trip.Pump = *req.Pump
    }
    if req.Rent != nil {
        trip.Rent = *req.Rent
    }
    if req.Advance != nil {
        trip.Advance = *req.Advance
    }
    if req.TripCost != nil {
        trip.TripCost = *req.TripCost
    }
    if req.Diesel != nil {
        trip.Diesel = *req.Diesel
    }
    if req.DieselPrice != nil {
        trip.DieselPrice = *req.DieselPrice
    }
    if req.ExtraCost != nil {
        trip.ExtraCost = *req.ExtraCost
    }

    // Recalculate
    trip.DieselTaka = trip.Diesel * trip.DieselPrice
    trip.Commission = trip.Rent * 0.1
    trip.Profit = trip.Rent - (trip.TripCost + trip.ExtraCost + trip.DieselTaka + trip.Commission)

    if err := s.repo.Update(trip); err != nil {
        return fmt.Errorf("failed to update own vehicle trip: %w", err)
    }

    return nil
}

func (s *ownVehicleTripService) Delete(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("own vehicle trip not found")
        }
        return err
    }

    if err := s.repo.Delete(id); err != nil {
        return fmt.Errorf("failed to delete own vehicle trip: %w", err)
    }

    return nil
}

func (s *ownVehicleTripService) GetByVehicleAndDateRange(vehicleNo, startDate, endDate string) ([]OwnVehicleTrip, error) {
    if vehicleNo == "" {
        return nil, errors.New("vehicle_no is required")
    }

    // Query by date field (string comparison, format: YYYY-MM-DD)
    trips, err := s.repo.GetByVehicleAndDateRange(vehicleNo, startDate, endDate)
    if err != nil {
        return nil, fmt.Errorf("failed to get trips by vehicle and date range: %w", err)
    }

    return trips, nil
}