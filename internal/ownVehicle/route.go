package ownvehicle

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, h *Handler) {
    v1 := router.Group("/api/v1")
    {
        // Build OwnVehicle view from Trip data by driver and date range
        v1.GET("/own-vehicles/by-driver", h.GetByDriverAndDate)
    }
}