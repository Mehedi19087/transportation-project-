package vehicle

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, vehicleHandler *VehicleHandler) {
    v1 := router.Group("/api/v1")
    {
        v1.POST("/vehicle", vehicleHandler.CreateVehicle)
        v1.GET("/vehicle/:id", vehicleHandler.GetVehicle)
        v1.PUT("/vehicle/:id", vehicleHandler.UpdateVehicle)
        v1.DELETE("/vehicle/:id", vehicleHandler.DeleteVehicle)
        v1.GET("/vehicle", vehicleHandler.GetAllVehicle)
    }
}