package dailysidecash

import (
    "time"

    "gorm.io/gorm"
)

type DailySideCash struct {
    ID               uint           `json:"id" gorm:"primaryKey"`
    Date             time.Time      `json:"date" gorm:"type:date;not null;index"`
    ProductID        uint           `json:"product_id" gorm:"index"`
    Cash             float64        `json:"cash"`
    RemainingBalance float64        `json:"remaining_balance"`
    ManagerName      string         `json:"manager_name"`
    TripCost         float64        `json:"trip_cost"`
    OtherCost        float64        `json:"other_cost"`
    OtherCostDetails string         `json:"other_cost_details"`
    CreatedAt        time.Time      `json:"created_at"`
    UpdatedAt        time.Time      `json:"updated_at"`
    DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

func (DailySideCash) TableName() string {
    return "daily_side_cash"
}