package bill 

import "time"

type CreateProductReq struct {
	 Name string `json:"name" binding:"required"`
	 CompanyID uint `json:"company_id" binding:"required"`
     ManagerName *string `json:"manager_name"`
	 Alt *int `json:"alt"`
	 Vat *int `json:"vat"`
}

type UpdateProductReq struct {
	 Name *string `json:"name"`
     ManagerName *string `json:"manager_name"`
	 Alt *int `json:"alt"`
	 Vat *int `json:"vat"`
}

type UpdateProductTripFieldsReq struct {
    TripFields []string `json:"trip_fields" binding:"required"`
}

type UpdateProductBillFieldsReq struct {
    BillFields []string `json:"bill_fields" binding:"required"`
}

type CreateBillReq struct {
    Category              *string  `json:"category"`
    ProductID             uint     `json:"product_id" binding:"required"`
    VehicleNo             *string  `json:"vehicle_no"`
    CustomerName          *string  `json:"customer_name"`
    ChallanNo             *string  `json:"challan_no"`
    DistributorName       *string  `json:"distributor_name"`
    DealerName            *string  `json:"dealer_name"`
    DriverName            *string  `json:"driver_name"`
    FromLocation          *string  `json:"from_location"`
    Destination           *string  `json:"destination"`
    Product               *string  `json:"product"`
    Portfolio             *string  `json:"portfolio"`
    Goods                 *string  `json:"goods"`
    Quantity              *float64 `json:"quantity"`
    BikeQty               *int     `json:"bike_qty"`
    VehicleSize           *string  `json:"vehicle_size"`
    Status                *string  `json:"status"`
    UnloadCharge          *float64 `json:"unload_charge"`
    VehicleRentWithVatTax *float64 `json:"vehicle_rent_with_vat_tax"`
    VehicleRent           *float64 `json:"vehicle_rent"`
    Dropping              *float64 `json:"dropping"`
    Alt5                  *float64 `json:"alt5"`
    Vat10                 *float64 `json:"vat10"`
    TotalRate             *float64 `json:"total_rate"`
    Advance               *float64 `json:"advance"`
    Due                   *float64 `json:"due"`
    Total                 *float64 `json:"total"`
    Profit                *float64 `json:"profit"`
    BodyFare              *float64 `json:"body_fare"`
    FuelCost              *float64 `json:"fuel_cost"`
    Amount                *float64 `json:"amount"`
    TotalAmount           *float64 `json:"total_amount"`
    DoNumber              *string  `json:"do_number"`
    CoNumber              *string  `json:"co_number"`
}

type UpdateBillReq struct {
    Category              *string  `json:"category"`
    ProductID             *uint    `json:"product_id"`
    VehicleNo             *string  `json:"vehicle_no"`
    CustomerName          *string  `json:"customer_name"`
    ChallanNo             *string  `json:"challan_no"`
    DistributorName       *string  `json:"distributor_name"`
    DealerName            *string  `json:"dealer_name"`
    DriverName            *string  `json:"driver_name"`
    FromLocation          *string  `json:"from_location"`
    Destination           *string  `json:"destination"`
    Product               *string  `json:"product"`
    Portfolio             *string  `json:"portfolio"`
    Goods                 *string  `json:"goods"`
    Quantity              *float64 `json:"quantity"`
    BikeQty               *int     `json:"bike_qty"`
    VehicleSize           *string  `json:"vehicle_size"`
    Status                *string  `json:"status"`
    UnloadCharge          *float64 `json:"unload_charge"`
    VehicleRentWithVatTax *float64 `json:"vehicle_rent_with_vat_tax"`
    VehicleRent           *float64 `json:"vehicle_rent"`
    Dropping              *float64 `json:"dropping"`
    Alt5                  *float64 `json:"alt5"`
    Vat10                 *float64 `json:"vat10"`
    TotalRate             *float64 `json:"total_rate"`
    Advance               *float64 `json:"advance"`
    Due                   *float64 `json:"due"`
    Total                 *float64 `json:"total"`
    Profit                *float64 `json:"profit"`
    BodyFare              *float64 `json:"body_fare"`
    FuelCost              *float64 `json:"fuel_cost"`
    Amount                *float64 `json:"amount"`
    TotalAmount           *float64 `json:"total_amount"`
    DoNumber              *string  `json:"do_number"`
    CoNumber              *string  `json:"co_number"`
}


type ProductSummary struct {
       ID   uint   `json:"id"`
       Name string `json:"name"`
       ManagerName string `json:"manager_name"`
       CompanyID   uint `json:"company_id"`
}

type CreateBillStatusReq struct {
	Date        time.Time `json:"date"`
	CompanyName string    `json:"company_name" binding:"required"`
	BillAmount  float64   `json:"bill_amount" binding:"required"`
	VatStatus   string    `json:"vat_status"`
	Status      string    `json:"status"`
}

type UpdateBillStatusReq struct {
	Date        *time.Time `json:"date"`
	CompanyName *string    `json:"company_name"`
	BillAmount  *float64   `json:"bill_amount"`
	VatStatus   *string    `json:"vat_status"`
	Status      *string    `json:"status"`
}