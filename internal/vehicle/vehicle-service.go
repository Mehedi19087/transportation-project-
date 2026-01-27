package vehicle

import (
    "errors"
    "fmt"
    "transportation/internal/utils"

    "gorm.io/gorm"
)

type VehicleService interface {
    CreateVehicle(req *CreateVehicle) error
    GetVehicle(id uint) (*Vehicle, error)
    UpdateVehicle(id uint, req *UpdateVehicle) error
    DeleteVehicle(id uint) error
    GetAllVehicle(cursorStr string, limit int) ([]Vehicle, string, error)
}

type vehicleService struct {
    repo VehicleRepo
}

func NewVehicleService(repo VehicleRepo) VehicleService {
    return &vehicleService{repo: repo}
}

func (s *vehicleService) CreateVehicle(req *CreateVehicle) error {
    if req.VehicleName == "" {
        return errors.New("vehicle name is required")
    }
    if req.Category == "" {
        return errors.New("category is required")
    }
    if req.VehicleNo == "" {
        return errors.New("vehicle number is required")
    }

    vehicle := &Vehicle{
        DriverID:         req.DriverID,
        DriverName:       req.DriverName,
        DriverContact:    req.DriverContact,
        Brand:            req.Brand,
        Model:            req.Model,
        Year:             req.Year,
        Mileage:          req.Mileage,
        FuelCapacity:     req.FuelCapacity,
        RegistrationDate: req.RegistrationDate,
        InsuranceExpiry:  req.InsuranceExpiry,
        FitnessExpiry:    req.FitnessExpiry,
        TaxTokenExpiry:   req.TaxTokenExpiry,
        PermitExpiry:     req.PermitExpiry,
        ImageURL:         req.ImageURL,
        VehicleName:      req.VehicleName,
        Category:         req.Category,
        Size:             req.Size,
        VehicleNo:        req.VehicleNo,
        Status:           req.Status,
        JoiningDate:      req.JoiningDate,
    }

    if vehicle.Status == "" {
        vehicle.Status = "active"
    }

    err := s.repo.Create(vehicle)
    if err != nil {
        return fmt.Errorf("failed to create vehicle: %w", err)
    }
    return nil
}

func (s *vehicleService) GetVehicle(id uint) (*Vehicle, error) {
    res, err := s.repo.Get(id)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *vehicleService) UpdateVehicle(id uint, req *UpdateVehicle) error {
    res, err := s.repo.Get(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return errors.New("vehicle data is missing")
        }
        return err
    }
if req.DriverID != nil {
        res.DriverID = req.DriverID
    }
    if req.DriverName != nil {
        res.DriverName = *req.DriverName
    }
    if req.DriverContact != nil {
        res.DriverContact = *req.DriverContact
    }
    if req.Brand != nil {
        res.Brand = *req.Brand
    }
    if req.Model != nil {
        res.Model = *req.Model
    }
    if req.Year != nil {
        res.Year = req.Year
    }
    if req.Mileage != nil {
        res.Mileage = req.Mileage
    }
    if req.FuelCapacity != nil {
        res.FuelCapacity = req.FuelCapacity
    }
	if req.RegistrationDate != nil {
        res.RegistrationDate = req.RegistrationDate
    }
    if req.InsuranceExpiry != nil {
        res.InsuranceExpiry = req.InsuranceExpiry
    }
    if req.FitnessExpiry != nil {
        res.FitnessExpiry = req.FitnessExpiry
    }
    if req.TaxTokenExpiry != nil {
        res.TaxTokenExpiry = req.TaxTokenExpiry
    }
    if req.PermitExpiry != nil {
        res.PermitExpiry = req.PermitExpiry
    }
    if req.ImageURL != nil {
        res.ImageURL = req.ImageURL
    }
    if req.VehicleName != nil {
        res.VehicleName = *req.VehicleName
    }
    if req.Category != nil {
        res.Category = *req.Category
    }

	if req.Size != nil {
        res.Size = *req.Size
    }
    if req.VehicleNo != nil {
        res.VehicleNo = *req.VehicleNo
    }
    if req.Status != nil {
        res.Status = *req.Status
    }
    if req.JoiningDate != nil {
        res.JoiningDate = *req.JoiningDate
    }

    if err := s.repo.Update(res); err != nil {
        return err
    }
    return nil
}

func (s *vehicleService) DeleteVehicle(id uint) error {
    err := s.repo.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s *vehicleService) GetAllVehicle(cursorStr string, limit int) ([]Vehicle, string, error) {
    if limit < 1 || limit > 100{
        limit = 10
    }

    cursor, err := utils.DecodeCursor(cursorStr)
    if err != nil {
        return nil, "", fmt.Errorf("invalid cursor: %w", err)
    }

    vehicles, err := s.repo.GetAll(cursor, limit)
    if err != nil {
        return nil, "", fmt.Errorf("list vehicles: %w", err)
    }

    var nextCursor string
    if len(vehicles) > limit {
        lastItem := vehicles[limit-1]
        nextCursor = utils.EncodeCursor(lastItem.CreatedAt, lastItem.ID)
        vehicles = vehicles[:limit]
    }

    return vehicles, nextCursor, nil
}