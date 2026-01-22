package calan

import (
	"time"

	"gorm.io/gorm"
)

type Calan struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Date        *string        `json:"date"`
	VehicleNo   *string        `json:"vehicle_no"`
	Destination *string        `json:"destination"`
	Amount      *float64       `json:"amount"`
	TripID      *uint          `json:"trip_id" gorm:"index"`
	OutsideTripID *uint        `json:"outside_trip_id" gorm:"index"`
	Status      string         `json:"status" gorm:"default:'unpaid'"` // paid / unpaid
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
