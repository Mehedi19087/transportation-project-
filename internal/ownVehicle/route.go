package ownvehicle

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, h *Handler) {
    v1 := router.Group("/api/v1")
    {
        // Build OwnVehicle view from Trip data by driver and date range
        v1.GET("/own-vehicles/by-driver", h.GetByDriverAndDate)
    }
}


func SetupRoutess(router *gin.Engine, handler *OwnVehicleTripHandler) {
     v1 := router.Group("/api/v1")
     {
          v1.POST("/own-vehicle-trips", handler.CreateOwnVehicleTrip)
          v1.GET("/own-vehicle-trips/:id", handler.GetOwnVehicleTrip)
          v1.GET("/own-vehicle-trips", handler.GetAllOwnVehicleTrips)
          v1.PUT("/own-vehicle-trips/:id", handler.UpdateOwnVehicleTrip)
          v1.DELETE("/own-vehicle-trips/:id", handler.DeleteOwnVehicleTrip)
          v1.GET("/own-vehicle-trips/by-driver", handler.GetOwnVehicleTripsByDriver)
     }
}
