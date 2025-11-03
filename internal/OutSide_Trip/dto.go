package outsidetrip

type OutSideTripReq struct {
    LoadPoint     string  `json:"load_point" binding:"required"`
    UnloadPoint   string  `json:"unload_point"`
    Rent          float64 `json:"rent"`
    Advance       float64 `json:"advance"`
    TripCost      float64 `json:"trip_cost"`
    Diesel        float64 `json:"diesel"`
    ExtraCost     float64 `json:"extra_cost"`
    DieselTaka    float64 `json:"diesel_taka"`
    Pamp          string  `json:"pamp"`
    Commission    float64 `json:"commission"`
    Month         string  `json:"month"`
    VehicleName   string  `json:"vehicle_name"`
    VehicleNumber string  `json:"vehicle_number"`
    DriverName    string  `json:"driver_name"`
    DriverPhone   string  `json:"driver_phone"`
}

type OutSideTripUpdateReq struct {
    LoadPoint     string  `json:"load_point"`
    UnloadPoint   string  `json:"unload_point"`
    Rent          float64 `json:"rent"`
    Advance       float64 `json:"advance"`
    TripCost      float64 `json:"trip_cost"`
    Diesel        float64 `json:"diesel"`
    ExtraCost     float64 `json:"extra_cost"`
    DieselTaka    float64 `json:"diesel_taka"`
    Pamp          string  `json:"pamp"`
    Commission    float64 `json:"commission"`
    Month         string  `json:"month"`
    VehicleName   string  `json:"vehicle_name"`
    VehicleNumber string  `json:"vehicle_number"`
    DriverName    string  `json:"driver_name"`
    DriverPhone   string  `json:"driver_phone"`
}

type VehicleMonthlySummary struct {
    VehicleNumber    string  `json:"vehicle_number"`
    Month            string  `json:"month"`
    TotalRent        float64 `json:"total_rent"`
    TotalAdvance     float64 `json:"total_advance"`
    TotalTripCost    float64 `json:"total_trip_cost"`
    TotalDiesel      float64 `json:"total_diesel"`
    TotalExtraCost   float64 `json:"total_extra_cost"`
    TotalDieselTaka  float64 `json:"total_diesel_taka"`
    TotalCommission  float64 `json:"total_commission"`
    TripCount        int64   `json:"trip_count"`
}