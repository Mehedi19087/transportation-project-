package routepricing

import "time"



type RoutePricing struct {
	 ID   uint `json:"id" gorm:"primaryKey"`
	 CustomerName string `json:"customer_name" gorm:"not null"`
	 Size int `json:"size" gorm:"not null"`
	 LoadPoint string `json:"load_point" gorm:"not null"`
	 UnloadPoint string `json:"unload_point" gorm:"not null"`
	 Rate int `json:"rate" gorm:"not null"`
	 Weight float64 `json:"weight"`
	 Notes  string `json:"notes"`
	 Status string `json:"status"`
	 VehicleCategory string `json:"vehicle_category"`
	 CreatedAt time.Time `json:"created_at"`
	 UpdatedAt time.Time `json:"updated_at"`
}
