package customer

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, customerHandler *CustomerHanlder) {
	 v1:=router.Group("/api/v1")
	 v1.POST("/customer",customerHandler.CreateCustomer)
	 v1.PUT("/customer/:id",customerHandler.UpdateCustomer)
	 v1.GET("/customer",customerHandler.GetAllCustomer)
	 v1.GET("/customer/:id",customerHandler.GetCustomer)
	 v1.DELETE("/customer/:id",customerHandler.DeleteCustomer)
}