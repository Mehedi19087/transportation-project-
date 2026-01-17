package routepricing

import "time"


type RoutePricing struct {
       ID          uint      `json:"id" gorm:"primaryKey"`
       ProductID   uint      `json:"product_id"`
       DealerName  string    `json:"dealer_name" gorm:"not null"`
       Destination string    `json:"destination" gorm:"not null"`
       Rate        float64   `json:"rate" gorm:"not null"`
       CreatedAt   time.Time `json:"created_at"`
       UpdatedAt   time.Time `json:"updated_at"`
}
