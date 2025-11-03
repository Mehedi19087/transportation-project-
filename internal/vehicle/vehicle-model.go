package vehicle

import "time"

type Vehicle struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    DriverID    *uint     `json:"driver_id"`

    DriverName  string    `json:"driver_name"`
	DriverContact    string     `json:"driver_contact"`
    Brand            string     `json:"brand"`
    Model            string     `json:"model"`
    Year             *int       `json:"year"`
    Mileage          *int       `json:"mileage"`
    FuelCapacity     *int       `json:"fuel_capacity"`         // frontend key uses 'fuleCapacity'
    RegistrationDate *time.Time `json:"registration_date"`
    InsuranceExpiry  *time.Time `json:"insurance_expiry"`
    FitnessExpiry    *time.Time `json:"fitness_expiry"`
    TaxTokenExpiry   *time.Time `json:"tax_token_expiry"`
    PermitExpiry     *time.Time `json:"permit_expiry"`
    ImageURL         *string    `json:"image_url"` 


    VehicleName string    `json:"vehicle_name" gorm:"not null"`
    Category    string    `json:"category" gorm:"not null"`
    Size        string    `json:"size"`
    VehicleNo   string    `json:"vehicle_no" gorm:"not null"`
    Status      string    `json:"status" gorm:"default:'active'"`
    JoiningDate time.Time `json:"joining_date"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}