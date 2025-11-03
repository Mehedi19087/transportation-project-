package trip

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, h *Handler) {
    v1 := router.Group("/api/v1")
    {
        v1.POST("/trips", h.CreateTrip)
        v1.GET("/trips/:id", h.GetTrip)
        v1.PUT("/trips/:id", h.UpdateTrip)
        v1.DELETE("/trips/:id", h.DeleteTrip)
        v1.GET("/products/:id/trips", h.GetProductTrips)
    }
}