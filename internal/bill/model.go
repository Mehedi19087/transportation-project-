package bill

import (
	"time"

	"github.com/lib/pq"
)

type Product struct {
	 ID uint `json:"id" gorm:"primaryKey"`
	 Name string `json:"name" gorm:"not null"`
	 CompanyID uint `json:"company_id" gorm:"index"`
	 Alt *int `json:"alt"`
	 Vat *int `json:"vat"`
     TripFields pq.StringArray `json:"trip_fields" gorm:"type:text[]"`
     BillFields pq.StringArray `json:"bill_fields" gorm:"type:text[]"`

	 CreatedAt time.Time `json:"created_at"`
}

type Bill struct {
    ID                    uint      `json:"id" gorm:"primaryKey"`
    Category              *string    `json:"category"`
    ProductID             uint      `json:"product_id" gorm:"not null; index"`
    VehicleNo             *string    `json:"vehicle_no"`
    CustomerName          *string    `json:"customer_name"`
    ChallanNo             *string    `json:"challan_no"`
    DistributorName       *string    `json:"distributor_name"`
    DealerName            *string    `json:"dealer_name"`
    DriverName            *string    `json:"driver_name"`
    FromLocation          *string    `json:"from"`
    Destination           *string    `json:"destination"`
    Product               *string    `json:"product"`
    Portfolio             *string    `json:"portfolio"`
    Goods                 *string    `json:"goods"`
    Quantity              *float64   `json:"quantity"`
    BikeQty               *int       `json:"bike_qty"`
    VehicleSize           *string    `json:"vehicle_size"`
    Status                *string    `json:"status"`
    UnloadCharge          *float64   `json:"unload_charge"`
    VehicleRentWithVATTax *float64   `json:"vehicle_rent_with_vat_tax"`
    VehicleRent           *float64   `json:"vehicle_rent"`
    Dropping              *float64   `json:"dropping"`
    Alt5                  *float64   `json:"alt5"`
    Vat10                 *float64   `json:"vat10"`
    TotalRate             *float64   `json:"total_rate"`
    Advance               *float64   `json:"advance"`
    Due                   *float64   `json:"due"`
    Total                 *float64   `json:"total"`
    Profit                *float64   `json:"profit"`
    BodyFare              *float64   `json:"body_fare"`
    FuelCost              *float64   `json:"fuel_cost"`
    Amount                *float64   `json:"amount"`
    TotalAmount           *float64   `json:"total_amount"`
    DoNumber              *string    `json:"do_number"`
    CoNumber              *string    `json:"co_number"`
    CreatedAt             time.Time `json:"created_at"`
}