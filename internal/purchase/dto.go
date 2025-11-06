package purchase

import "time"

type CreatePurchase struct {
    SupplierName  string    `json:"supplier_name" form:"supplier_name" binding:"required"`
    OrderDate     time.Time `json:"order_date" form:"order_date" binding:"required"`
    DeliveryDate  time.Time `json:"delivery_date" form:"delivery_date" binding:"required"`
    Items         string    `json:"items" form:"items" binding:"required"`
    Quantity      string    `json:"quantity" form:"quantity" binding:"required"`
    Category      string    `json:"category" form:"category" binding:"required"`
    TotalAmount   float64   `json:"total_amount" form:"total_amount" binding:"required"`
    VehicleNo     string    `json:"vehicle_no" form:"vehicle_no" binding:"required"`
    Notes         *string   `json:"notes" form:"notes"`
    PaymentStatus string    `json:"payment_status" form:"payment_status" binding:"required"` // Paid, Pending, Overdue
    BillImage     *string   `json:"bill_image" form:"bill_image"`                                              // Will be set if file uploaded
}

type UpdatePurchase struct {
    SupplierName  string    `json:"supplier_name" form:"supplier_name"`
    OrderDate     time.Time `json:"order_date" form:"order_date"`
    DeliveryDate  time.Time `json:"delivery_date" form:"delivery_date"`
    Items         string    `json:"items" form:"items"`
    Quantity      string    `json:"quantity" form:"quantity"`
    Category      string    `json:"category" form:"category"`
    TotalAmount   float64   `json:"total_amount" form:"total_amount"`
    VehicleNo     string    `json:"vehicle_no" form:"vehicle_no"`
    Notes         *string   `json:"notes" form:"notes"`
    PaymentStatus string    `json:"payment_status" form:"payment_status"`
    BillImage     *string   `json:"bill_image" form:"bill_image"` // Will be set if file uploaded
}