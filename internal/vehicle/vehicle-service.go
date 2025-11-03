package vehicle

import (
    "errors"
    "fmt"
    "gorm.io/gorm"
)

type VehicleService interface {
    CreateVehicle(req *CreateVehicle) error
    GetVehicle(id uint) (*Vehicle, error)
    UpdateVehicle(id uint, req *UpdateVehicle) error
    DeleteVehicle(id uint) error
    GetAllVehicle(page, pageSize int) ([]Vehicle, int64, error)
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

func (s *vehicleService) GetAllVehicle(page, pageSize int) ([]Vehicle, int64, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }
    offset := (page - 1) * pageSize

    vehicles, total, err := s.repo.GetAll(offset, pageSize)
    if err != nil {
        return nil, 0, fmt.Errorf("list vehicles: %w", err)
    }
    return vehicles, total, nil
}