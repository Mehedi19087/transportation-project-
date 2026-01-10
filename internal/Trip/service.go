// package trip

// import (
//     "errors"
//     "time"

//     "gorm.io/gorm"
// )

// type Service interface {
//     CreateTrip(req *CreateTripReq) error
//     GetTrip(id uint) (*Trip, error)
//     GetProductTrips(productID uint, page, pageSize int) ([]Trip, int64, error)
//     UpdateTrip(id uint, req *UpdateTripReq) error
//     DeleteTrip(id uint) error
// }

// type service struct {
//     repo Repository
// }

// func NewService(r Repository) Service {
//     return &service{repo: r}
// }

// func (s *service) CreateTrip(req *CreateTripReq) error {
//     if req.ProductID == 0 {
//         return errors.New("product_id is required")
//     }
//     item := &Trip{
//         ProductID:     req.ProductID,
//         BrandName:     req.BrandName,
//         Category:      req.Category,
//         Date:          req.Date,
//         TripType:      req.TripType,
//         TripNo:        req.TripNo,
//         InvoiceNo:     req.InvoiceNo,
//         VehicleName:   req.VehicleName,
//         VehicleNo:     req.VehicleNo,
//         EngineNo:      req.EngineNo,
//         ChassisNo:     req.ChassisNo,
//         DriverName:    req.DriverName,
//         DriverMobile:  req.DriverMobile,
//         HelperName:    req.HelperName,
//         LoadPoint:     req.LoadPoint,
//         UnloadPoint:   req.UnloadPoint,
//         Destination:   req.Destination,
//         Route:         req.Route,
//         District:      req.District,
//         Quantity:      req.Quantity,
//         UnitPrice:     req.UnitPrice,
//         Cash:          req.Cash,
//         Advance:       req.Advance,
//         Due:           req.Due,
//         BillNo:        req.BillNo,
//         BillDate:      req.BillDate,
//         PaymentType:   req.PaymentType,
//         TruckSize:     req.TruckSize,
//         Weight:        req.Weight,
//         FuelType:      req.FuelType,
//         FuelCost:      req.FuelCost,
//         TransportType: req.TransportType,
//         Remarks:       req.Remarks,
//         Status:        req.Status,
//         CreatedBy:     req.CreatedBy,
//         ApprovedBy:    req.ApprovedBy,
//         CreatedAt:     time.Now(),
//         UpdatedAt:     time.Now(),
//     }

//     return s.repo.Create(item)
// }

// func (s *service) GetTrip(id uint) (*Trip, error) {
//     return s.repo.Get(id)
// }

// func (s *service) GetProductTrips(productID uint, page, pageSize int) ([]Trip, int64, error) {
//     if productID == 0 {
//         return nil, 0, errors.New("invalid product id")
//     }
//     if page < 1 {
//         page = 1
//     }
//     if pageSize < 1 {
//         pageSize = 10
//     }
//     offset := (page - 1) * pageSize
//     return s.repo.GetByProduct(productID, offset, pageSize)
// }

// func (s *service) UpdateTrip(id uint, req *UpdateTripReq) error {
//     item, err := s.repo.Get(id)
//     if err != nil {
//         if err == gorm.ErrRecordNotFound {
//             return errors.New("trip not found")
//         }
//         return err
//     }

//     if req.ProductID != nil {
//         item.ProductID = *req.ProductID
//     }
//     if req.BrandName != nil {
//         item.BrandName = req.BrandName
//     }
//     if req.Category != nil {
//         item.Category = req.Category
//     }
//     if req.Date != nil {
//         item.Date = req.Date
//     }
//     if req.TripType != nil {
//         item.TripType = req.TripType
//     }
//     if req.TripNo != nil {
//         item.TripNo = req.TripNo
//     }
//     if req.InvoiceNo != nil {
//         item.InvoiceNo = req.InvoiceNo
//     }
//     if req.VehicleName != nil {
//         item.VehicleName = req.VehicleName
//     }
//     if req.VehicleNo != nil {
//         item.VehicleNo = req.VehicleNo
//     }
//     if req.EngineNo != nil {
//         item.EngineNo = req.EngineNo
//     }
//     if req.ChassisNo != nil {
//         item.ChassisNo = req.ChassisNo
//     }
//     if req.DriverName != nil {
//         item.DriverName = req.DriverName
//     }
//     if req.DriverMobile != nil {
//         item.DriverMobile = req.DriverMobile
//     }
//     if req.HelperName != nil {
//         item.HelperName = req.HelperName
//     }
//     if req.LoadPoint != nil {
//         item.LoadPoint = req.LoadPoint
//     }
//     if req.UnloadPoint != nil {
//         item.UnloadPoint = req.UnloadPoint
//     }
//     if req.Destination != nil {
//         item.Destination = req.Destination
//     }
//     if req.Route != nil {
//         item.Route = req.Route
//     }
//     if req.District != nil {
//         item.District = req.District
//     }
//     if req.Quantity != nil {
//         item.Quantity = req.Quantity
//     }
//     if req.UnitPrice != nil {
//         item.UnitPrice = req.UnitPrice
//     }
//     if req.TotalRate != nil {
//         item.TotalRate = req.TotalRate
//     }
//     if req.Cash != nil {
//         item.Cash = req.Cash
//     }
//     if req.Advance != nil {
//         item.Advance = req.Advance
//     }
//     if req.Due != nil {
//         item.Due = req.Due
//     }
//     if req.BillNo != nil {
//         item.BillNo = req.BillNo
//     }
//     if req.BillDate != nil {
//         item.BillDate = req.BillDate
//     }
//     if req.PaymentType != nil {
//         item.PaymentType = req.PaymentType
//     }
//     if req.TruckSize != nil {
//         item.TruckSize = req.TruckSize
//     }
//     if req.Weight != nil {
//         item.Weight = req.Weight
//     }
//     if req.FuelType != nil {
//         item.FuelType = req.FuelType
//     }
//     if req.FuelCost != nil {
//         item.FuelCost = req.FuelCost
//     }
//     if req.TransportType != nil {
//         item.TransportType = req.TransportType
//     }
//     if req.Remarks != nil {
//         item.Remarks = req.Remarks
//     }
//     if req.Status != nil {
//         item.Status = req.Status
//     }
//     if req.CreatedBy != nil {
//         item.CreatedBy = req.CreatedBy
//     }
//     if req.ApprovedBy != nil {
//         item.ApprovedBy = req.ApprovedBy
//     }
//     item.UpdatedAt = time.Now()

//     return s.repo.Update(item)
// }

// func (s *service) DeleteTrip(id uint) error {
//     _, err := s.repo.Get(id)
//     if err != nil {
//         if err == gorm.ErrRecordNotFound {
//             return errors.New("trip not found")
//         }
//         return err
//     }
//     return s.repo.Delete(id)
// }

package trip

import (
	"errors"
	"fmt"
	"time"
	"log"

	"gorm.io/gorm"
	"transportation/internal/ownVehicle"
	
	"transportation/internal/routePricing"
)

type Service interface {
	CreateTrip(req *CreateTripReq) error
	GetTrip(id uint) (*Trip, error)
	GetProductTrips(productID uint, page, pageSize int) ([]Trip, int64, error)
	UpdateTrip(id uint, req *UpdateTripReq) error
	DeleteTrip(id uint) error
}

type service struct {
    repo                Repository
    routePricingService routepricing.RoutePricingService
    ownVehicleTripRepo  ownvehicle.OwnVehicleTripRepository
}

func NewService(r Repository, rps routepricing.RoutePricingService, ovr ownvehicle.OwnVehicleTripRepository) Service {
    return &service{
        repo:                r,
        routePricingService: rps,
        ownVehicleTripRepo:  ovr,
    }
}

// Helper method to calculate fields
func (s *service) calculateFields(item *Trip) error {
	// Derive total rate from advance + due (user does not supply total rate)
	advance := 0.0
	if item.Advance != nil {
		advance = *item.Advance
	}
	due := 0.0
	if item.Due != nil {
		due = *item.Due
	}
	totalRate := advance + due
	item.TotalRate = &totalRate

	// Fetch route rate (if available) for VAT/profit calculation
	routeRate := 0.0
	if item.LoadPoint != nil && item.UnloadPoint != nil && *item.LoadPoint != "" && *item.UnloadPoint != "" {
		rate, err := s.routePricingService.GetRate(*item.Dealer, *item.UnloadPoint)
		if err == nil {
			routeRate = float64(rate)
			vat := routeRate * 0.20
			item.Vat10 = &vat
		}
	}

	// Profit = totalRate + (20% of route rate) - unit price
	// unitPrice := 0.0
	// if item.UnitPrice != nil {
	// 	unitPrice = *item.UnitPrice
	// }
	profit := routeRate - (totalRate + (routeRate * 0.20))
	item.Profit = &profit

	// Set default status
	if item.Status == nil || *item.Status == "" {
		status := "Pending"
		item.Status = &status
	}

	return nil
}

func (s *service) CreateTrip(req *CreateTripReq) error {
    if req.ProductID == 0 {
        return errors.New("product_id is required")
    }

    log.Printf("DEBUG: CreateTrip - Dealer: '%s', UnloadPoint: '%s'", req.Dealer, req.UnloadPoint)
    rate, err := s.routePricingService.GetRate(req.Dealer, req.UnloadPoint)
    if err != nil {
        log.Printf("ERROR: GetRate failed for Dealer: '%s', UnloadPoint: '%s': %v", req.Dealer, req.UnloadPoint, err)
        return fmt.Errorf("route pricing is missing for dealer '%s' and destination '%s'", req.Dealer, req.UnloadPoint)
    }

    unitPrice := float64(rate)
    vat := unitPrice * 0.2

    item := &Trip{
        ProductID:     req.ProductID,
        BrandName:     &req.BrandName,
        Category:      &req.Category,
        Date:          &req.Date,
        TripType:      &req.TripType,
        TripNo:        &req.TripNo,
        InvoiceNo:     &req.InvoiceNo,
        VehicleName:   &req.VehicleName,
        VehicleNo:     &req.VehicleNo,
        EngineNo:      &req.EngineNo,
        ChassisNo:     &req.ChassisNo,
        DriverName:    &req.DriverName,
        DriverMobile:  &req.DriverMobile,
        HelperName:    &req.HelperName,
        LoadPoint:     &req.LoadPoint,
        UnloadPoint:   &req.UnloadPoint,
        Destination:   &req.Destination,
        Route:         &req.Route,
        District:      &req.District,
        Quantity:      &req.Quantity,
        UnitPrice:     &unitPrice,
        Cash:          &req.Cash,
        Advance:       &req.Advance,
        Due:           &req.Due,
        BillNo:        &req.BillNo,
        BillDate:      &req.BillDate,
        PaymentType:   &req.PaymentType,
        TruckSize:     &req.TruckSize,
        Weight:        &req.Weight,
        FuelType:      &req.FuelType,
        FuelCost:      &req.FuelCost,
        TransportType: &req.TransportType,
        Remarks:       &req.Remarks,
        Status:        &req.Status,
        CreatedBy:     &req.CreatedBy,
        ApprovedBy:    &req.ApprovedBy,
        Dealer:        &req.Dealer,
        Rent:          &req.Rent,
        Alt5:          &req.Alt5,
        Vat10:         &vat,
        CreatedAt:     time.Now().UTC(),
        UpdatedAt:     time.Now().UTC(),
    }

    // Calculate all dependent fields
    if err := s.calculateFields(item); err != nil {
        return err
    }

    // Start transaction
    db := s.repo.GetDB()
    tx := db.Begin()
    if tx.Error != nil {
        return fmt.Errorf("failed to start transaction: %w", tx.Error)
    }

    // 1. Create Trip
    if err := s.repo.CreateWithTx(tx, item); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to create trip: %w", err)
    }

    // 2. Create OwnVehicleTrip (only if VehicleNo exists)
    if req.VehicleNo != "" {
        // Check using transaction
        var existingTrip ownvehicle.OwnVehicleTrip
        err := tx.Where("trip_id = ?", item.ID).First(&existingTrip).Error

        if err == gorm.ErrRecordNotFound {
            // Not found - create new
            rent := unitPrice
            commission := rent * 0.1

            ownVehicleTrip := &ownvehicle.OwnVehicleTrip{
                Date:        req.Date,
                VehicleNo:   req.VehicleNo,
                LoadPoint:   req.LoadPoint,
                UnloadPoint: req.UnloadPoint,
                Rent:        rent,
                Advance:     req.Advance,
                TripCost:    0,
                Diesel:      0,
                DieselPrice: 0,
                DieselTaka:  0,
                ExtraCost:   0,
                Commission:  commission,
                TripID:      &item.ID,
            }

            if err := s.ownVehicleTripRepo.CreateWithTx(tx, ownVehicleTrip); err != nil {
                tx.Rollback()
                return fmt.Errorf("failed to create own vehicle trip: %w", err)
            }
        } else if err != nil {
            // Unexpected DB error (connection, timeout, etc.)
            tx.Rollback()
            return fmt.Errorf("failed to check own vehicle trip: %w", err)
        }
        // If no error (record exists), do nothing
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}

// ...existing GetTrip, GetProductTrips, UpdateTrip, DeleteTrip code (keep as is)...









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

// ...existing code...

func (s *service) UpdateTrip(id uint, req *UpdateTripReq) error {
    item, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("trip not found")
        }
        return err
    }

    // Start transaction
    db := s.repo.GetDB()
    tx := db.Begin()
    if tx.Error != nil {
        return fmt.Errorf("failed to start transaction: %w", tx.Error)
    }

    // ...existing field updates...
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
    if req.Dealer != nil {
        item.Dealer = req.Dealer
    }
    if req.Rent != nil {
        item.Rent = req.Rent
    }
    if req.Alt5 != nil {
        item.Alt5 = req.Alt5
    }
    if req.Vat10 != nil {
        item.Vat10 = req.Vat10
    }

    item.UpdatedAt = time.Now()

    // Recalculate all dependent fields
    if err := s.calculateFields(item); err != nil {
        tx.Rollback()
        return err
    }

    // Update Trip
    if err := s.repo.UpdateWithTx(tx, item); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to update trip: %w", err)
    }

    // Handle OwnVehicleTrip (only if VehicleNo exists)
    vehicleNo := ""
    if item.VehicleNo != nil {
        vehicleNo = *item.VehicleNo
    }

    if vehicleNo != "" {
        // Check if OwnVehicleTrip exists
        var ownTrip ownvehicle.OwnVehicleTrip
        err = tx.Where("trip_id = ?", item.ID).First(&ownTrip).Error

        if err == gorm.ErrRecordNotFound {
            // *** NOT FOUND - CREATE NEW ***
            rent := 0.0
            if item.UnitPrice != nil {
                rent = *item.UnitPrice
            }
            advance := 0.0
            if item.Advance != nil {
                advance = *item.Advance
            }
            date := ""
            if item.Date != nil {
                date = *item.Date
            }
            loadPoint := ""
            if item.LoadPoint != nil {
                loadPoint = *item.LoadPoint
            }
            unloadPoint := ""
            if item.UnloadPoint != nil {
                unloadPoint = *item.UnloadPoint
            }

            newOwnTrip := &ownvehicle.OwnVehicleTrip{
                Date:        date,
                VehicleNo:   vehicleNo,
                LoadPoint:   loadPoint,
                UnloadPoint: unloadPoint,
                Rent:        rent,
                Advance:     advance,
                TripCost:    0,
                Diesel:      0,
                DieselPrice: 0,
                DieselTaka:  0,
                ExtraCost:   0,
                Commission:  rent * 0.1,
                TripID:      &item.ID,
            }

            if err := s.ownVehicleTripRepo.CreateWithTx(tx, newOwnTrip); err != nil {
                tx.Rollback()
                return fmt.Errorf("failed to create own vehicle trip: %w", err)
            }

        } else if err != nil {
            // Some other DB error
            tx.Rollback()
            return fmt.Errorf("failed to check own vehicle trip: %w", err)

        } else {
            // *** FOUND - UPDATE EXISTING ***
            if item.Date != nil {
                ownTrip.Date = *item.Date
            }
            ownTrip.VehicleNo = vehicleNo
            if item.LoadPoint != nil {
                ownTrip.LoadPoint = *item.LoadPoint
            }
            if item.UnloadPoint != nil {
                ownTrip.UnloadPoint = *item.UnloadPoint
            }
            if item.Advance != nil {
                ownTrip.Advance = *item.Advance
            }
            if item.UnitPrice != nil {
                ownTrip.Rent = *item.UnitPrice
            }

            // Recalculate derived fields
            ownTrip.DieselTaka = ownTrip.Diesel * ownTrip.DieselPrice
            ownTrip.Commission = ownTrip.Rent * 0.1

            if err := s.ownVehicleTripRepo.UpdateWithTx(tx, &ownTrip); err != nil {
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

// ...existing code...

// ...existing code...

func (s *service) DeleteTrip(id uint) error {
    _, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("trip not found")
        }
        return err
    }

    // Start transaction
    db := s.repo.GetDB()
    tx := db.Begin()
    if tx.Error != nil {
        return fmt.Errorf("failed to start transaction: %w", tx.Error)
    }

    // 1. Delete OwnVehicleTrip first (if exists)
    if err := s.ownVehicleTripRepo.DeleteByTripIDWithTx(tx, id); err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete own vehicle trip: %w", err)
    }

    // 2. Delete Trip
    if err := tx.Delete(&Trip{}, id).Error; err != nil {
        tx.Rollback()
        return fmt.Errorf("failed to delete trip: %w", err)
    }

    // Commit transaction
    if err := tx.Commit().Error; err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}
