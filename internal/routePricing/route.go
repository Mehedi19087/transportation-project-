package routepricing

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, routePricingHandler *RoutePricingHandler) {
    v1 := router.Group("/api/v1")
    {
        v1.POST("/route-pricing", routePricingHandler.CreateRoutePricing)
        v1.GET("/route-pricing/:id", routePricingHandler.GetRoutePricing)
        v1.PUT("/route-pricing/:id", routePricingHandler.UpdateRoutePricing)
        v1.DELETE("/route-pricing/:id", routePricingHandler.DeleteRoutePricing)
        v1.GET("/route-pricing", routePricingHandler.GetAllRoutePricing)
        v1.GET("/rates",routePricingHandler.GetRate)
    }
}