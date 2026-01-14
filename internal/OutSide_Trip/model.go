package outsidetrip

import (
	"time"

	"gorm.io/gorm"
)


type OutSideTrip struct {
	 ID uint `json:"id" gorm:"primaryKey"`
	 Date string `json:"date" gorm:"index"`
	 LoadPoint string `json:"load_point"`
	 UnloadPoint string `json:"unload_point"`
	 Rent float64 `json:"rent"`
	 Advance float64 `json:"advance"`
	 TripCost float64 `json:"trip_cost"`
	 Diesel float64 `json:"diesel"`
	 ExtraCost float64 `json:"extra_cost"`
	 DieselPrice float64 `json:"diesel_price" gorm:"not null"`
	 DieselTaka float64 `json:"diesel_taka"`
	 Pamp string `json:"pamp"`
	 Commission float64 `json:"commission"`
	 Month string `gorm:"index:idx_vehicle_month" json:"month"`
	 VehicleName string `json:"vehicle_name"`
	 VehicleNumber string `json:"vehicle_number" gorm:"index:idx_vehicle_month"`
	 DriverName string `json:"driver_name"`
	 DriverPhone string `json:"driver_phone"`
	 ProductID  uint `json:"product_id" gorm:"not null;index"`
	 Due        float64 `json:"due"`
	 DueStatus  string  `json:"due_status" gorm:"default:unpaid"`
	 CreatedAt time.Time `json:"created_at"`
	 UpdatedAt time.Time `json:"updated_at"`
	 DeletedAt gorm.DeletedAt `json:"_" gorm:"index"`
}
func(OutSideTrip) TableName() string {
	 return "outside_trips"
}