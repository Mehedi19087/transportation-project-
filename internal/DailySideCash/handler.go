package dailysidecash

  import (
      "fmt"
      "net/http"
      "strconv"

      "github.com/gin-gonic/gin"
  )

  type Handler struct {
      svc Service
  }

  func NewHandler(svc Service) *Handler {
      return &Handler{svc: svc}
  }

  func (h *Handler) Create(c *gin.Context) {
      productIDVal, exists := c.Get("product_id")
      if !exists {
          c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
          return
      }
      productID := productIDVal.(uint)

      var req CreateDailySideCashDTO
      if err := c.ShouldBindJSON(&req); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }
      
      fmt.Printf("DEBUG: Received Date: %q\n", req.Date)

      rec, err := h.svc.Create(&req, productID)
      if err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }

      c.JSON(http.StatusCreated, gin.H{"message": "created", "data": rec})
  }

  func (h *Handler) Update(c *gin.Context) {
      idStr := c.Param("id")
      id, err := strconv.ParseUint(idStr, 10, 32)
      if err != nil || id == 0 {
          c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
          return
      }

      var req UpdateDailySideCashDTO
      if err := c.ShouldBindJSON(&req); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }

      if err := h.svc.Update(uint(id), &req); err != nil {
          if err.Error() == "daily side cash not found" {
              c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
              return
          }
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }

      c.JSON(http.StatusOK, gin.H{"message": "updated"})
  }

  func (h *Handler) Get(c *gin.Context) {
      idStr := c.Param("id")
      id, err := strconv.ParseUint(idStr, 10, 32)
      if err != nil || id == 0 {
          c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
          return
      }

      rec, err := h.svc.Get(uint(id))
      if err != nil {
          if err.Error() == "daily side cash not found" {
              c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
              return
          }
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }

      c.JSON(http.StatusOK, gin.H{"data": rec})
  }

  func (h *Handler) GetAll(c *gin.Context) {
      productIDVal, exists := c.Get("product_id")
      if !exists {
          c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
          return
      }
      productID := productIDVal.(uint)

      page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
      pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
      if page < 1 {
          page = 1
      }
      if pageSize < 1 {
          pageSize = 10
      }

      records, total, err := h.svc.GetAll(productID, page, pageSize)
      if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }

      c.JSON(http.StatusOK, gin.H{
          "data": records,
          "meta": gin.H{
              "page":      page,
              "page_size": pageSize,
              "total":     total,
          },
      })
  }

  func (h *Handler) Delete(c *gin.Context) {
      idStr := c.Param("id")
      id, err := strconv.ParseUint(idStr, 10, 32)
      if err != nil || id == 0 {
          c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
          return
      }

      if err := h.svc.Delete(uint(id)); err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }

      c.JSON(http.StatusOK, gin.H{"message": "deleted"})
  }

  func (h *Handler) GetByDate(c *gin.Context) {
      productIDVal, exists := c.Get("product_id")
      if !exists {
          c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
          return
      }
      productID := productIDVal.(uint)

      date := c.Query("date")
      if date == "" {
          c.JSON(http.StatusBadRequest, gin.H{"error": "date is required (DD-MM-YYYY)"})
          return
      }

      rec, err := h.svc.GetByDate(productID, date)
      if err != nil {
          if err.Error() == "daily side cash not found" {
              c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
              return
          }
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }

      c.JSON(http.StatusOK, gin.H{"data": rec})
  }
