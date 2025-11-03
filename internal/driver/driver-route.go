package driver

import "github.com/gin-gonic/gin"


func SetupRoutes(router *gin.Engine, driverHandler *DriverHandler){
	v1:= router.Group("/api/v1")
	{
		 v1.POST("/driver", driverHandler.CreateDriver)
		 v1.GET("/driver/:id", driverHandler.GetDriver)
         v1.PUT("/driver/:id", driverHandler.UpdateDriver)
         v1.DELETE("/driver/:id", driverHandler.DeleteDriver)
         v1.GET("/driver", driverHandler.GetAllDriver)
	}
}