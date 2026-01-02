package ownvehicle

import (
	"time"
	"transportation/internal/utils"

	"gorm.io/gorm"
)


type OwnVehicle struct {
	 ID uint `json:"id" gorm:"primaryKey"`
	 LoadPoint string `json:"load_point"`
	 UnloadPoint string `json:"unload_point"`
	 Rent float64 `json:"rent"`
	 Advance float64 `json:"advance"`
	 TripCost float64 `json:"trip_cost"`
	 Diesel float64 `json:"diesel"`
	 ExtraCost float64 `json:"extra_cost"`
	 DieselTaka float64 `json:"diesel_taka"`
	 Pamp string `json:"pamp"`
	 Commission float64 `json:"commission"`
	 TripID   uint `json:"trip_id" gorm:"not null;uniqueIndex"`
	 CreatedAt time.Time `json:"created_at"`
	 UpdatedAt time.Time `json:"updated_at"`
	 DeletedAt gorm.DeletedAt `json:"-"`
}
func (OwnVehicle) TableName() string {
    return "own_vehicles"
}


type OwnVehicleTrip struct {
	 ID  uint `json:"id" gorm:"primaryKey"`
	 VehicleNo string `json:"vehicle_no" gorm:"index; not null"`
	 DriverName string `json:"driver_name" gorm:"index; not null"`
	 LoadPoint  string `json:"load_poin" gorm:"not null"`
	 UnloadPoint string `json:"unload_point" gorm:"not null"`
	 Rent        float64 `json:"rent" gorm:"not null"`
	 Advance     float64 `json:"advance" gorm:"not null"`
	 TripCost    float64  `json:"trip_cost" gorm:"not null"`
	 Diesel       float64   `json:"diesel" gorm:"not null"`
	 DieselPrice float64    `json:"diesel_price" gorm:"not null"`
	 DieselTaka  float64     `json:"diesel_taka" gorm:"not null"`
	 ExtraCost   float64      `json:"extra_cost" gorm:"not null"`
	 Commission  float64      `json:"commission" gorm:"not null"`
	 Profit      float64      `json:"profit" gorm:"not null"`
	 Pump        string        `json:"pump" gorm:"not null"`
	 CreatedAt   utils.BDTime    `json:"created_at"`
	 UpdatedAt   utils.BDTime    `json:"updated_at"`
	 DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

func(OwnVehicleTrip) TableName() string {
	 return "own_vehicle_trips"
}