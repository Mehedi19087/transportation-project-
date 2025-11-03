package driver

import (
	"time"
)

type CreateDriver struct {
    DriverName  string    `json:"driver_name" form:"driver_name" binding:"required"`
    Mobile      string    `json:"mobile" form:"mobile"`
    Address     string    `json:"address" form:"address"`
    Emergency   string    `json:"emergency" form:"emergency"`
    License     string    `json:"license" form:"license"`
    Expired     time.Time `json:"expired" form:"expired"`
    NID         string    `json:"nid" form:"nid"`
    JoiningDate time.Time `json:"joining_date" form:"joining_date"`
    Status      string    `json:"status" form:"status"`
    ImageURL    *string   `json:"image_url"`
}

type UpdateDriver struct {
    DriverName  string    `json:"driver_name" form:"driver_name"`
    Mobile      string    `json:"mobile" form:"mobile"`
    Address     string    `json:"address" form:"address"`
    Emergency   string    `json:"emergency" form:"emergency"`
    License     string    `json:"license" form:"license"`
    Expired     time.Time `json:"expired" form:"expired"`
    NID         string    `json:"nid" form:"nid"`
    JoiningDate time.Time `json:"joining_date" form:"joining_date"`
    Status      string    `json:"status" form:"status"`
    ImageURL    *string   `json:"image_url"`
}