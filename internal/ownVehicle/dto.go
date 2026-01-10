package ownvehicle

// VehicleDateQuery is used for searching history/ledger by vehicle and date range
type VehicleDateQuery struct {
    VehicleNumber string `form:"vehicle_number" binding:"required"`
    StartDate     string `form:"start_date" binding:"required"` // expected "YYYY-MM-DD"
    EndDate       string `form:"end_date" binding:"required"`   // expected "YYYY-MM-DD"
}

type HistoryRequest struct {
    VehicleNumber string `form:"vehicle_number" binding:"required"`
    StartDate     string `form:"start_date" binding:"required"` // expected "YYYY-MM-DD"
    EndDate       string `form:"end_date" binding:"required"`   // expected "YYYY-MM-DD"
}

type CreateOwnVehicleTripDTO struct {
    Date         string  `json:"date" binding:"required"` // YYYY-MM-DD
    VehicleNo    string  `json:"vehicle_no" binding:"required"`
    LoadPoint    string  `json:"load_point" binding:"required"`
    UnloadPoint  string  `json:"unload_point" binding:"required"`
    Rent         float64 `json:"rent" binding:"required"`
    Advance      float64 `json:"advance"`
    Pump         string  `json:"pump"`
    TripCost     float64 `json:"trip_cost"`
    Diesel       float64 `json:"diesel"`
    DieselPrice  float64 `json:"diesel_price"`
    ExtraCost    float64 `json:"extra_cost"`
    TripID       *uint   `json:"trip_id"`
    OutsideTripID *uint  `json:"outside_trip_id"`
}

type UpdateOwnVehicleTripDTO struct {
    Date         *string  `json:"date"`
    VehicleNo    *string  `json:"vehicle_no"`
    LoadPoint    *string  `json:"load_point"`
    UnloadPoint  *string  `json:"unload_point"`
    Rent         *float64 `json:"rent"`
    Advance      *float64 `json:"advance"`
    Pump         *string  `json:"pump"`
    TripCost     *float64 `json:"trip_cost"`
    Diesel       *float64 `json:"diesel"`
    DieselPrice  *float64 `json:"diesel_price"`
    ExtraCost    *float64 `json:"extra_cost"`
    Commission   *float64 `json:"commission"`
}


type DriverDateQuery struct {
    DriverName string `form:"driver_name" binding:"required"`
    StartDate  string `form:"start_date" binding:"required"`
    EndDate    string `form:"end_date" binding:"required"`
}