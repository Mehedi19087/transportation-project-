package employee

import (
	"transportation/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, employeeHandler *EmployeeHandler) {
    v1 := router.Group("/api/v1")
    v1.Use(auth.AuthMiddleware())
    v1.POST("/employee", employeeHandler.CreateEmployee)
    v1.GET("/employee/:id", employeeHandler.GetEmployee)
    v1.PUT("/employee/:id", employeeHandler.UpdateEmployee)
    v1.DELETE("/employee/:id", employeeHandler.DeleteEmployee)
    v1.GET("/employee", employeeHandler.GetAllEmployee)
}