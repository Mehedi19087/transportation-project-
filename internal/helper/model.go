package helper

import (
	"time"

	"gorm.io/gorm"
)

type Helper struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name" gorm:"not null"`
	Mobile          string         `json:"mobile" gorm:"not null"`
	Emergency       string         `json:"emergency"`
	Address         string         `json:"address"`
	Salary          string         `json:"salary"`
	AssignedVehicle string         `json:"assignedVehicle"`
	JoiningDate     *time.Time     `json:"joiningDate"` // pointer so it can be null
	Experience      string         `json:"experience"`
	NID             string         `json:"nid" gorm:"column:nid"`
	Status          string         `json:"status" gorm:"default:Active"`
	Image           string         `json:"image"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}
