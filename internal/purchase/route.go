package purchase

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, h *PurchaseHandler) {
    v1 := router.Group("/api/v1")
    {
        v1.POST("/purchase", h.CreatePurchase)
        v1.GET("/purchase/:id", h.GetPurchase)
        v1.PUT("/purchase/:id", h.UpdatePurchase)
        v1.DELETE("/purchase/:id", h.DeletePurchase)
        v1.GET("/purchase", h.GetAllPurchase)
    }
}