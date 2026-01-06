package dailysidecash

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
