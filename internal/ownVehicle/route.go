package ownvehicle

import (
    "transportation/internal/auth"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, h *Handler) {
    v1 := router.Group("/api/v1")
    v1.Use(auth.AuthMiddleware())
    {
        // Build OwnVehicle view from Trip data by driver and date range
        v1.GET("/own-vehicles/by-driver", h.GetByDriverAndDate)
    }
}

func SetupOwnVehicleTripRoutes(router *gin.Engine, handler *OwnVehicleTripHandler) {
    v1 := router.Group("/api/v1")
    v1.Use(auth.AuthMiddleware())
    {
        // Static routes FIRST (before :id)
        v1.GET("/own-vehicle-trips/by-vehicle", handler.GetOwnVehicleTripsByVehicle)
        v1.GET("/own-vehicle-trips", handler.GetAllOwnVehicleTrips)
        v1.POST("/own-vehicle-trips", handler.CreateOwnVehicleTrip)
        
        // Dynamic routes LAST
        v1.GET("/own-vehicle-trips/:id", handler.GetOwnVehicleTrip)
        v1.PUT("/own-vehicle-trips/:id", handler.UpdateOwnVehicleTrip)
        v1.DELETE("/own-vehicle-trips/:id", handler.DeleteOwnVehicleTrip)
    }
}