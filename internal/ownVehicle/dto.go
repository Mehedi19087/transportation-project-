package ownvehicle  

type DriverDateQuery struct {
    DriverName string `form:"driver_name" binding:"required"`
    StartDate  string `form:"start_date" binding:"required"` // expected "YYYY-MM-DD"
    EndDate    string `form:"end_date" binding:"required"`   // expected "YYYY-MM-DD"
}

type CreateOwnVehicleTripDTO struct {
    VehicleNo    string  `json:"vehicle_no" binding:"required"`
    DriverName   string  `json:"driver_name" binding:"required"`
    LoadPoint    string  `json:"load_point" binding:"required"`
    UnloadPoint  string  `json:"unload_point" binding:"required"`
    Pump         string  `json:"pump" binding:"required"`
    Rent         float64 `json:"rent" binding:"required"`
    Advance      float64 `json:"advance" binding:"required"`
    TripCost     float64 `json:"trip_cost" binding:"required"`
    Diesel       float64 `json:"diesel" binding:"required"`
    DieselPrice  float64 `json:"diesel_price" binding:"required"`
    ExtraCost    float64 `json:"extra_cost" binding:"required"`
}

type UpdateOwnVehicleTripDTO struct {
    VehicleNo    *string  `json:"vehicle_no"`
    DriverName   *string  `json:"driver_name"`
    LoadPoint    *string  `json:"load_point"`
    UnloadPoint  *string  `json:"unload_point"`
    Pump         *string  `json:"pump"`
    Rent         *float64 `json:"rent"`
    Advance      *float64 `json:"advance"`
    TripCost     *float64 `json:"trip_cost"`
    Diesel       *float64 `json:"diesel"`
    DieselPrice  *float64 `json:"diesel_price"`
    ExtraCost    *float64 `json:"extra_cost"`

}