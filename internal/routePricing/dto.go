package routepricing

type RoutePricingReq struct {
    CustomerName string  `json:"customer_name" binding:"required"`
    Size         int     `json:"size" binding:"required"`
    LoadPoint    string  `json:"load_point" binding:"required"`
    UnloadPoint  string  `json:"unload_point" binding:"required"`
    Rate         int     `json:"rate" binding:"required"`
    VehicleCategory string `json:"vehicle_category"`
    Weight       float64 `json:"weight"`
    Notes        string  `json:"notes"`
    Status       string  `json:"status"`
}

type RoutePricingUpdateReq struct {
    CustomerName string  `json:"customer_name"`
    Size         int     `json:"size"`
    LoadPoint    string  `json:"load_point"`
    UnloadPoint  string  `json:"unload_point"`
    VehicleCategory string `json:"vehicle_category"`
    Rate         int     `json:"rate"`
    Weight       float64 `json:"weight"`
    Notes        string  `json:"notes"`
    Status       string  `json:"status"`
}

type RateRequest struct {
    LoadPoint   string `json:"load_point" binding:"required"`
    UnloadPoint string `json:"unload_point" binding:"required"`
}