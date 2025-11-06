package trip 

import (
    "encoding/json"
    "time"
)
type CreateTripReq struct {
    ProductID     uint     `json:"product_id" binding:"required"`
    BrandName     *string  `json:"brand_name"`
    Category      *string  `json:"category"`
    Date          *string  `json:"date"`
    TripType      *string  `json:"trip_type"`
    TripNo        *string  `json:"trip_no"`
    InvoiceNo     *string  `json:"invoice_no"`
    VehicleName   *string  `json:"vehicle_name"`
    VehicleNo     *string  `json:"vehicle_no"`
    EngineNo      *string  `json:"engine_no"`
    ChassisNo     *string  `json:"chassis_no"`
    DriverName    *string  `json:"driver_name"`
    DriverMobile  *string  `json:"driver_mobile"`
    HelperName    *string  `json:"helper_name"`
    LoadPoint     *string  `json:"load_point"`
    UnloadPoint   *string  `json:"unload_point"`
    Destination   *string  `json:"destination"`
    Route         *string  `json:"route"`
    District      *string  `json:"district"`
    Quantity      *float64 `json:"quantity"`
    UnitPrice     *float64 `json:"unit_price"`
    TotalRate     *float64 `json:"total_rate"`
    Cash          *float64 `json:"cash"`
    Advance       *float64 `json:"advance"`
    Due           *float64 `json:"due"`
    BillNo        *string  `json:"bill_no"`
    BillDate      *string  `json:"bill_date"`
    PaymentType   *string  `json:"payment_type"`
    TruckSize     *string  `json:"truck_size"`
    Weight        *float64 `json:"weight"`
    FuelType      *string  `json:"fuel_type"`
    FuelCost      *float64 `json:"fuel_cost"`
    TransportType *string  `json:"transport_type"`
    Remarks       *string  `json:"remarks"`
    Status        *string  `json:"status"`
    CreatedBy     *string  `json:"created_by"`
    ApprovedBy    *string  `json:"approved_by"`
}

type UpdateTripReq struct {
    ProductID     *uint    `json:"product_id"`
    BrandName     *string  `json:"brand_name"`
    Category      *string  `json:"category"`
    Date          *string  `json:"date"`
    TripType      *string  `json:"trip_type"`
    TripNo        *string  `json:"trip_no"`
    InvoiceNo     *string  `json:"invoice_no"`
    VehicleName   *string  `json:"vehicle_name"`
    VehicleNo     *string  `json:"vehicle_no"`
    EngineNo      *string  `json:"engine_no"`
    ChassisNo     *string  `json:"chassis_no"`
    DriverName    *string  `json:"driver_name"`
    DriverMobile  *string  `json:"driver_mobile"`
    HelperName    *string  `json:"helper_name"`
    LoadPoint     *string  `json:"load_point"`
    UnloadPoint   *string  `json:"unload_point"`
    Destination   *string  `json:"destination"`
    Route         *string  `json:"route"`
    District      *string  `json:"district"`
    Quantity      *float64 `json:"quantity"`
    UnitPrice     *float64 `json:"unit_price"`
    TotalRate     *float64 `json:"total_rate"`
    Cash          *float64 `json:"cash"`
    Advance       *float64 `json:"advance"`
    Due           *float64 `json:"due"`
    BillNo        *string  `json:"bill_no"`
    BillDate      *string  `json:"bill_date"`
    PaymentType   *string  `json:"payment_type"`
    TruckSize     *string  `json:"truck_size"`
    Weight        *float64 `json:"weight"`
    FuelType      *string  `json:"fuel_type"`
    FuelCost      *float64 `json:"fuel_cost"`
    TransportType *string  `json:"transport_type"`
    Remarks       *string  `json:"remarks"`
    Status        *string  `json:"status"`
    CreatedBy     *string  `json:"created_by"`
    ApprovedBy    *string  `json:"approved_by"`
}


func (t Trip) MarshalJSON() ([]byte, error) {
    loc, err := time.LoadLocation("Asia/Dhaka")
    if err != nil {
        loc = time.FixedZone("BDT", 6*60*60)
    }
    type Alias Trip
    return json.Marshal(&struct {
        CreatedAt string `json:"created_at"`
        UpdatedAt string `json:"updated_at"`
        *Alias
    }{
        CreatedAt: t.CreatedAt.In(loc).Format("2006-01-02 15:04:05"),
        UpdatedAt: t.UpdatedAt.In(loc).Format("2006-01-02 15:04:05"),
        Alias:     (*Alias)(&t),
    })
}