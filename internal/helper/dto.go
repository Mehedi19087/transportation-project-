package helper


import (
	"time"
)

// Create input
type CreateHelperDTO struct {
	Name            string     `json:"name" binding:"required"`
	Mobile          string     `json:"mobile" binding:"required"`
	Emergency       string     `json:"emergency"`
	Address         string     `json:"address"`
	Salary          string     `json:"salary"`
	AssignedVehicle string     `json:"assignedVehicle"`
	JoiningDate     *time.Time `json:"joiningDate"` // ISO8601 string in JSON
	Experience      string     `json:"experience"`
	NID             string     `json:"nid"`
	Status          string     `json:"status"` // default to Active if empty
	Image           string     `json:"image"`
}

// Update input (all optional except ID in path)
type UpdateHelperDTO struct {
	Name            *string    `json:"name"`
	Mobile          *string    `json:"mobile"`
	Emergency       *string    `json:"emergency"`
	Address         *string    `json:"address"`
	Salary          *string    `json:"salary"`
	AssignedVehicle *string    `json:"assignedVehicle"`
	JoiningDate     *time.Time `json:"joiningDate"`
	Experience      *string    `json:"experience"`
	NID             *string    `json:"nid"`
	Status          *string    `json:"status"`
	Image           *string    `json:"image"`
}

// Response
type HelperResponse struct {
	ID              uint       `json:"id"`
	Name            string     `json:"name"`
	Mobile          string     `json:"mobile"`
	Emergency       string     `json:"emergency"`
	Address         string     `json:"address"`
	Salary          string     `json:"salary"`
	AssignedVehicle string     `json:"assignedVehicle"`
	JoiningDate     *time.Time `json:"joiningDate"`
	Experience      string     `json:"experience"`
	NID             string     `json:"nid"`
	Status          string     `json:"status"`
	Image           string     `json:"image"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
