package ownvehicle

import (
    "fmt"
    "time"
)

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

    // 3) Map trips -> OwnVehicle view
    result := make([]OwnVehicle, 0, len(trips))
    for _, t := range trips {
        lp := derefStr(t.LoadPoint)
        up := derefStr(t.UnloadPoint)
        adv := derefF64(t.Advance)

        rate, rateErr := s.repo.GetRateByLocations(lp, up)
        if rateErr != nil {
            rate = 0 // no configured rate
        }

        ov := OwnVehicle{
			TripID: t.ID,
            LoadPoint:   lp,
            UnloadPoint: up,
            Rent:        rate,
            Advance:     adv,
            Commission:  rate * 0.20, // 20%
            // TripCost, Diesel, ExtraCost, DieselTaka, Pamp left as zero-values
        }
	_= s.repo.CreateOwnVehicle(&ov)
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