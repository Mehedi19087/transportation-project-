package auth

import (
    "github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, handler *AuthHandler) {
    auth := router.Group("api/v1/auth")
    {
        auth.POST("/register", handler.CreateUser)
        auth.POST("/login", handler.Login)
        
        // Protected routes
        protected := auth.Group("/")
        protected.Use(AuthMiddleware())
        {
            protected.PUT("/users/:id", handler.UpdateUser)
            
            // Admin only routes
            admin := protected.Group("/")
            admin.Use(RequireRole(RoleAdmin))
            {
                admin.GET("/pending", handler.GetPendingUsers)
            }
        }
    }
}