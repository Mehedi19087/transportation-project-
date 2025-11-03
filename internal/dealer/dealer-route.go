package dealer

import "github.com/gin-gonic/gin"


func SetupRoutes(router *gin.Engine, dealerhandler *DealerHanlder) {
	 v1:= router.Group("/api/v1")
	 v1.POST("/dealer",dealerhandler.CreateDealer)
	 v1.GET("/dealer/:id",dealerhandler.GetDealer)
	 v1.PUT("/dealer/:id",dealerhandler.UpdateDealer)
	 v1.DELETE("/dealer/:id",dealerhandler.DeleteDealer)
	 v1.GET("/dealer",dealerhandler.GetAllDealer)
}