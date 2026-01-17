package routepricing

type RoutePricingReq struct {
    ProductID   uint    `json:"product_id" binding:"required"`
    DealerName  string  `json:"dealer_name" binding:"required"`
    Destination string  `json:"destination" binding:"required"`
    Rate        float64 `json:"rate" binding:"required"`
}

type RoutePricingUpdateReq struct {
    ProductID   uint    `json:"product_id"`
    DealerName  string  `json:"dealer_name"`
    Destination string  `json:"destination"`
    Rate        float64 `json:"rate"`
}

type RateRequest struct {
    ProductID   uint   `json:"product_id" binding:"required"`
    DealerName  string `json:"dealer_name" binding:"required"`
    Destination string `json:"destination" binding:"required"`
}