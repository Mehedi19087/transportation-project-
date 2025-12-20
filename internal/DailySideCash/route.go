
  package dailysidecash

  import "github.com/gin-gonic/gin"

  func SetupRoutes(router *gin.Engine, h *Handler) {
      v1 := router.Group("/api/v1")
      {
          v1.POST("/daily-side-cash", h.Create)
          v1.GET("/daily-side-cash/:id", h.Get)
          v1.GET("/daily-side-cash", h.GetAll)          // paginated list
          v1.PUT("/daily-side-cash/:id", h.Update)
          v1.DELETE("/daily-side-cash/:id", h.Delete)
          v1.GET("/daily-side-cash/by-date", h.GetByDate) // expects ?date=YYYY-MM-DD
      }
  }