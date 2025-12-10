package helper

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, helperHandler *HelperHandler) {
    v1 := router.Group("/api/v1")
    {
        v1.POST("/helper", helperHandler.CreateHelper)
        v1.GET("/helper/:id", helperHandler.GetHelper)
        v1.PUT("/helper/:id", helperHandler.UpdateHelper)
        v1.DELETE("/helper/:id", helperHandler.DeleteHelper)
        v1.GET("/helpers", helperHandler.GetAllHelper)
    }
}