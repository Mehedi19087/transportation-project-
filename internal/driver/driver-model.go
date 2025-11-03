package driver

import (
	"time"
)

type Driver struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"not null"`
    Mobile      string    `json:"mobile" gorm:"unique;not null"`
	ImageUrl *string      `json:"image" gorm:"column:image_url"`
    Address     string    `json:"address"`
    Emergency   string    `json:"emergency"`
    License     string    `json:"license" gorm:"unique"`
    Expired     time.Time `json:"expired"`
    NID         string    `json:"nid" gorm:"column:nid;unique"`
    JoiningDate time.Time `json:"joining_date"`
    Status      string    `json:"status" gorm:"default:'active'"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Vehicle struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    DriverID     *uint     `json:"driver_id"`  // nullable FK, can be NULL if unassigned
    VehicleName  string    `json:"vehicle_name" gorm:"not null"`
    Category     string    `json:"category" gorm:"not null"`
    Size         string    `json:"size"`
    VehicleNo    string    `json:"vehicle_no" gorm:"not null"`
    Status       string    `json:"status" gorm:"default:'active'"`
    JoiningDate  time.Time `json:"joining_date"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}