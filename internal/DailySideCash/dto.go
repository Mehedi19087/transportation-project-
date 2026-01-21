package dailysidecash

import (
	"time"
	"transportation/internal/utils"
)

// CreateDailySideCashDTO is the request body for creating a daily side cash record.
// Expect "YYYY-MM-DD" for Date.
type CreateDailySideCashDTO struct {
	Date             string  `json:"date" binding:"required"`
	Cash             float64 `json:"cash"`
	RemainingBalance float64 `json:"remaining_balance"`
	TripCost         float64 `json:"trip_cost"`
	OtherCost        float64 `json:"other_cost"`
	OtherCostDetails string  `json:"other_cost_details"`
	ProductID        uint    `json:"product_id"`
}

// UpdateDailySideCashDTO is the request body for partial updates.
type UpdateDailySideCashDTO struct {
	Date             *string  `json:"date"`
	Cash             *float64 `json:"cash"`
	RemainingBalance *float64 `json:"remaining_balance"`
	TripCost         *float64 `json:"trip_cost"`
	OtherCost        *float64 `json:"other_cost"`
	OtherCostDetails *string  `json:"other_cost_details"`
}

// Response DTOs with BD time
type DailySideCashResponse struct {
	ID               uint            `json:"id"`
	Date             time.Time       `json:"date"` // Date is stored as date only, so it won't need timezone conversion
	ProductID        uint            `json:"product_id"`
	Cash             float64         `json:"cash"`
	WithoutRemaining float64         `json:"without_remaining"`
	RemainingBalance float64         `json:"remaining_balance"`
	ManagerName      string          `json:"manager_name"`
	TripCost         float64         `json:"trip_cost"`
	OtherCost        float64         `json:"other_cost"`
	OtherCostDetails string          `json:"other_cost_details"`
	CreatedAt        utils.BDTime    `json:"created_at"`
	UpdatedAt        utils.BDTime    `json:"updated_at"`
}
