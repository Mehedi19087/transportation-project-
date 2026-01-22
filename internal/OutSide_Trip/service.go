package outsidetrip

import (
	"errors"
	"fmt"
	"time"
	"transportation/internal/bill"
	"transportation/internal/calan"
	ownvehicle "transportation/internal/ownVehicle"

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
    repo                OutSideTripRepo
    ownVehicleTripRepo  ownvehicle.OwnVehicleTripRepository
    calanRepo           calan.CalanRepo
    billRepo            bill.ProductRepo
}

func NewOutSideTripService(repo OutSideTripRepo, ownoutSideTripRepo ownvehicle.OwnVehicleTripRepository, calanRepo calan.CalanRepo, billRepo bill.ProductRepo) OutSideTripService {
    return &outSideTripService {
        repo:               repo,
        ownVehicleTripRepo: ownoutSideTripRepo,
        calanRepo:          calanRepo,
        billRepo:           billRepo,
    }
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
        //ProductID:     req.ProductID,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }

    //start transaction 
    db := s.repo.GetDB()
    tx := db.Begin()
    if tx.Error != nil {
        return fmt.Errorf("failed to start transaction: %w", tx.Error)
    }

    if err := s.repo.CreateWithTx(tx, trip); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to create trip: %w", err)
    }

    // Create Bill
    status := "Pending"
    newBill := &bill.Bill{
        ProductID:     trip.ProductID,
        OutsideTripID: &trip.ID,
        Date:          &trip.Date,
        VehicleNo:     &trip.VehicleNumber,
        DriverName:    &trip.DriverName,
        FromLocation:  &trip.LoadPoint,
        Destination:   &trip.UnloadPoint,
        Amount:        &trip.Rent,
        TotalAmount:   &trip.Rent,
        Advance:       &trip.Advance,
        Due:           &trip.Due,
        Status:        &status,
        CreatedAt:     time.Now().UTC(),
    }
    if err := s.billRepo.CreateBillWithTx(tx, newBill); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to create bill: %w", err)
    }

    if req.LoadPoint != "" {
         var existing ownvehicle.OwnVehicleTrip 
         err := tx.Where("trip_id = ?",trip.ID).First(&existing).Error

         if err == gorm.ErrRecordNotFound {
            ownVehicleTrip := &ownvehicle.OwnVehicleTrip{
                Date:        req.Date,
                VehicleNo:   req.VehicleNumber,
                LoadPoint:   req.LoadPoint,
                UnloadPoint: req.UnloadPoint,
                Rent:        req.Rent,
                Advance:     req.Advance,
                TripCost:    req.TripCost,
                Diesel:      req.Diesel,
                DieselPrice: req.DieselPrice,
                DieselTaka:  req.Diesel * req.DieselPrice,
                ExtraCost:   req.ExtraCost,
                Commission:  req.Rent * 0.1,
                TripID:      &trip.ID,
            }
         
         if err := s.ownVehicleTripRepo.CreateWithTx(tx, ownVehicleTrip); err != nil {
                tx.Rollback()
                return fmt.Errorf("failed to create own vehicle trip: %w", err)
            }
        } else if err != nil { 
            tx.Rollback()
            return fmt.Errorf("failed to check own vehicle trip: %w", err)
        }
    }

    // Create Calan if Due > 0
    if req.Due > 0 {
        newCalan := &calan.Calan{
            Date:          &req.Date,
            VehicleNo:     &req.VehicleNumber,
            Destination:   &req.UnloadPoint,
            Amount:        &req.Due,
            OutsideTripID: &trip.ID,
            Status:        "unpaid",
            CreatedAt:     time.Now().UTC(),
        }
        if err := s.calanRepo.Create(tx, newCalan); err != nil {
            tx.Rollback()
            return fmt.Errorf("failed to create calan: %w", err)
        }
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    
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

    // Start transaction
    db := s.repo.GetDB()
    tx := db.Begin()
    if tx.Error != nil {
        return fmt.Errorf("failed to start transaction: %w", tx.Error)
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
    if req.DieselPrice != nil {
        res.DieselPrice = *req.DieselPrice
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

    if err := s.repo.UpdateWithTx(tx, res); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to update outside trip: %w", err)
    }

    // Update Bill
    var billItem bill.Bill
    if err := tx.Where("outside_trip_id = ?", res.ID).First(&billItem).Error; err == nil {
        billItem.Date = &res.Date
        billItem.VehicleNo = &res.VehicleNumber
        billItem.DriverName = &res.DriverName
        billItem.FromLocation = &res.LoadPoint
        billItem.Destination = &res.UnloadPoint
        billItem.Amount = &res.Rent
        billItem.TotalAmount = &res.Rent
        billItem.Advance = &res.Advance
        billItem.Due = &res.Due
        
        if err := tx.Save(&billItem).Error; err != nil {
             tx.Rollback()
             return fmt.Errorf("failed to update bill: %w", err)
        }
    }

    // Update Calan
    var calanItem calan.Calan
    if err := tx.Where("outside_trip_id = ?", res.ID).First(&calanItem).Error; err == nil {
        calanItem.Date = &res.Date
        calanItem.VehicleNo = &res.VehicleNumber
        calanItem.Destination = &res.UnloadPoint
        calanItem.Amount = &res.Due
        
        if err := tx.Save(&calanItem).Error; err != nil {
             tx.Rollback()
             return fmt.Errorf("failed to update calan: %w", err)
        }
    }

    // Update OwnVehicleTrip
    if res.LoadPoint != "" {
         var existing ownvehicle.OwnVehicleTrip 
         err := tx.Where("trip_id = ?", res.ID).First(&existing).Error

         if err == gorm.ErrRecordNotFound {
            // Create if not exists
            ownVehicleTrip := &ownvehicle.OwnVehicleTrip{
                Date:        res.Date,
                VehicleNo:   res.VehicleNumber,
                LoadPoint:   res.LoadPoint,
                UnloadPoint: res.UnloadPoint,
                Rent:        res.Rent,
                Advance:     res.Advance,
                TripCost:    res.TripCost,
                Diesel:      res.Diesel,
                DieselPrice: res.DieselPrice,
                DieselTaka:  res.Diesel * res.DieselPrice,
                ExtraCost:   res.ExtraCost,
                Commission:  res.Rent * 0.1,
                TripID:      &res.ID,
            }
         
            if err := s.ownVehicleTripRepo.CreateWithTx(tx, ownVehicleTrip); err != nil {
                tx.Rollback()
                return fmt.Errorf("failed to create own vehicle trip: %w", err)
            }
        } else if err != nil { 
            tx.Rollback()
            return fmt.Errorf("failed to check own vehicle trip: %w", err)
        } else {
            // Update existing
            existing.Date = res.Date
            existing.VehicleNo = res.VehicleNumber
            existing.LoadPoint = res.LoadPoint
            existing.UnloadPoint = res.UnloadPoint
            existing.Rent = res.Rent
            existing.Advance = res.Advance
            existing.TripCost = res.TripCost
            existing.Diesel = res.Diesel
            existing.DieselPrice = res.DieselPrice
            existing.DieselTaka = res.Diesel * res.DieselPrice
            existing.ExtraCost = res.ExtraCost
            existing.Commission = res.Rent * 0.1
            
            if err := s.ownVehicleTripRepo.UpdateWithTx(tx, &existing); err != nil {
                tx.Rollback()
                return fmt.Errorf("failed to update own vehicle trip: %w", err)
            }
        }
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}

func (s *outSideTripService) DeleteOutSideTrip(id uint) error {
    // Start transaction
    db := s.repo.GetDB()
    tx := db.Begin()
    if tx.Error != nil {
        return fmt.Errorf("failed to start transaction: %w", tx.Error)
    }

    // 1. Delete OwnVehicleTrip first (if exists)
    // We use the same 'trip_id' logic as in Create/Update
    if err := s.ownVehicleTripRepo.DeleteByTripIDWithTx(tx, id); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete own vehicle trip: %w", err)
    }

    // 2. Delete Associated Bill
    if err := s.billRepo.DeleteBillByOutsideTripID(tx, id); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete associated bill: %w", err)
    }

    // 3. Delete Associated Calan
    if err := s.calanRepo.DeleteByOutsideTripID(tx, id); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete associated calan: %w", err)
    }

    // 4. Delete OutSideTrip
    if err := s.repo.DeleteWithTx(tx, id); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete outside trip: %w", err)
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
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