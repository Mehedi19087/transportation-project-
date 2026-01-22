package calan

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, handler *CalanHandler) {
	group := router.Group("/api/calans")
	{
		group.GET("", handler.GetCalans)
        group.PATCH("/:id/status", handler.UpdateStatus)
	}
}
