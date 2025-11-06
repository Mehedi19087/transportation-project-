package ownvehicle

import (
	"time"

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