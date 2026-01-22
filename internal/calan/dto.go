package calan

import "time"

type CalanResponse struct {
	ID            uint      `json:"id"`
	Date          *string   `json:"date"`
	VehicleNo     *string   `json:"vehicle_no"`
	Destination   *string   `json:"destination"`
	Amount        *float64  `json:"amount"`
	TripID        *uint     `json:"trip_id,omitempty"`
	OutsideTripID *uint     `json:"outside_trip_id,omitempty"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type FilterCalanReq struct {
	Date      string `form:"date"`
	VehicleNo string `form:"vehicle_no"`
	Status    string `form:"status"` // "paid" or "unpaid"
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
}
