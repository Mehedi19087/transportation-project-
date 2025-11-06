package purchase

import (
    "time"

    "gorm.io/gorm"
)

type Purchase struct {
    ID            uint           `json:"id" gorm:"primaryKey"`
    SupplierName  *string         `json:"supplier_name"`
    OrderDate     *time.Time      `json:"order_date"`
    DeliveryDate  *time.Time      `json:"delivery_date"`
    Items         *string         `json:"items"`
    Quantity      *string         `json:"quantity"`
    Category      *string         `json:"category"`
    TotalAmount   *float64        `json:"total_amount"`
    VehicleNo     *string         `json:"vehicle_no"`
    Notes         *string        `json:"notes"`
    PaymentStatus *string         `json:"payment_status"` // e.g., Paid, Pending, Overdue
    BillImage     *string        `json:"bill_image"` // URL or base64
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `json:"-"`
}

func (Purchase) TableName() string {
    return "purchases"
}