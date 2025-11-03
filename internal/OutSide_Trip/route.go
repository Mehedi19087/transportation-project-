package outsidetrip

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, outsideTripHandler *OutSideTripHandler) {
    v1 := router.Group("/api/v1")
    v1.POST("/outside-trip", outsideTripHandler.CreateOutSideTrip)
    v1.GET("/outside-trip/:id", outsideTripHandler.GetOutSideTrip)
    v1.PUT("/outside-trip/:id", outsideTripHandler.UpdateOutSideTrip)
    v1.DELETE("/outside-trip/:id", outsideTripHandler.DeleteOutSideTrip)
    v1.GET("/outside-trip", outsideTripHandler.GetAllOutSideTrip)
	v1.GET("/vehicles/:vehicleNumber/months/:month/summary", outsideTripHandler.GetVehicleMonthlySummary)
}