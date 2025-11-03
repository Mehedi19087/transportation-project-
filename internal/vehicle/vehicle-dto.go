package vehicle

import "time"

type CreateVehicle struct {
    DriverID         *uint      `json:"driver_id" form:"driver_id"`
    DriverName       string     `json:"driver_name" form:"driver_name"`
    DriverContact    string     `json:"driver_contact" form:"driver_contact"`
    
    Brand            string     `json:"brand" form:"brand"`
    Model            string     `json:"model" form:"model"`
    Year             *int       `json:"year" form:"year"`
    Mileage          *int       `json:"mileage" form:"mileage"`
    FuelCapacity     *int       `json:"fuel_capacity" form:"fuel_capacity"`
    RegistrationDate *time.Time `json:"registration_date" form:"registration_date"`
    InsuranceExpiry  *time.Time `json:"insurance_expiry" form:"insurance_expiry"`
    FitnessExpiry    *time.Time `json:"fitness_expiry" form:"fitness_expiry"`
    TaxTokenExpiry   *time.Time `json:"tax_token_expiry" form:"tax_token_expiry"`
    PermitExpiry     *time.Time `json:"permit_expiry" form:"permit_expiry"`
    ImageURL         *string    `json:"image_url"`
    
    VehicleName      string     `json:"vehicle_name" form:"vehicle_name" binding:"required"`
    Category         string     `json:"category" form:"category" binding:"required"`
    Size             string     `json:"size" form:"size"`
    VehicleNo        string     `json:"vehicle_no" form:"vehicle_no" binding:"required"`
    Status           string     `json:"status" form:"status"`
    JoiningDate      time.Time  `json:"joining_date" form:"joining_date"`
}

type UpdateVehicle struct {
    DriverID         *uint      `json:"driver_id" form:"driver_id"`
    DriverName       *string    `json:"driver_name" form:"driver_name"`
    DriverContact    *string    `json:"driver_contact" form:"driver_contact"`
    
    Brand            *string    `json:"brand" form:"brand"`
    Model            *string    `json:"model" form:"model"`
    Year             *int       `json:"year" form:"year"`
    Mileage          *int       `json:"mileage" form:"mileage"`
    FuelCapacity     *int       `json:"fuel_capacity" form:"fuel_capacity"`
    RegistrationDate *time.Time `json:"registration_date" form:"registration_date"`
    InsuranceExpiry  *time.Time `json:"insurance_expiry" form:"insurance_expiry"`
    FitnessExpiry    *time.Time `json:"fitness_expiry" form:"fitness_expiry"`
    TaxTokenExpiry   *time.Time `json:"tax_token_expiry" form:"tax_token_expiry"`
    PermitExpiry     *time.Time `json:"permit_expiry" form:"permit_expiry"`
    ImageURL         *string    `json:"image_url"`
    
    VehicleName      *string    `json:"vehicle_name" form:"vehicle_name"`
    Category         *string    `json:"category" form:"category"`
    Size             *string    `json:"size" form:"size"`
    VehicleNo        *string    `json:"vehicle_no" form:"vehicle_no"`
    Status           *string    `json:"status" form:"status"`
    JoiningDate      *time.Time `json:"joining_date" form:"joining_date"`
}