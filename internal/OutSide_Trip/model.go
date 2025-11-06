package outsidetrip

import "time"


type OutSideTrip struct {
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
	 Month string `gorm:"index:idx_vehicle_month" json:"month"`
	 VehicleName string `json:"vehicle_name"`
	 VehicleNumber string `json:"vehicle_number" gorm:"index:idx_vehicle_month"`
	 DriverName string `json:"driver_name"`
	 DriverPhone string `json:"driver_phone"`
	 CreatedAt time.Time `json:"created_at"`
	 UpdatedAt time.Time `json:"updated_at"`
}
func(OutSideTrip) TableName() string {
	 return "outside_trips"
}