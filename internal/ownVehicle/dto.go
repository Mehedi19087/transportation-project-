package ownvehicle  

type DriverDateQuery struct {
    DriverName string `form:"driver_name" binding:"required"`
    StartDate  string `form:"start_date" binding:"required"` // expected "YYYY-MM-DD"
    EndDate    string `form:"end_date" binding:"required"`   // expected "YYYY-MM-DD"
}