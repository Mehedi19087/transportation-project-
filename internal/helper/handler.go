package helper

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type HelperHandler struct {
    service HelperService
}

func NewHelperHandler(service HelperService) *HelperHandler {
    return &HelperHandler{service: service}
}

func (h *HelperHandler) CreateHelper(c *gin.Context) {
    var req CreateHelperDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    helper, err := h.service.CreateHelper(&req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Helper created successfully",
        "helper":  helper,
    })
}

func (h *HelperHandler) GetHelper(c *gin.Context) {
    id := c.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid helper ID"})
        return
    }
    if idStr == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be zero"})
        return
    }

    helper, err := h.service.GetHelper(uint(idStr))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": helper})
}

func (h *HelperHandler) UpdateHelper(c *gin.Context) {
    id := c.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid helper ID"})
        return
    }
    if idStr == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be zero"})
        return
    }

    var req UpdateHelperDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.UpdateHelper(uint(idStr), &req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Helper updated successfully"})
}

func (h *HelperHandler) DeleteHelper(c *gin.Context) {
    id := c.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid helper ID"})
        return
    }
    if idStr == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be zero"})
        return
    }

    if err := h.service.DeleteHelper(uint(idStr)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Helper deleted successfully"})
}

func (h *HelperHandler) GetAllHelper(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")
    pageSizeStr := c.DefaultQuery("page_size", "10")

    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }

    pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil || pageSize < 1 {
        pageSize = 10
    }

    helpers, total, err := h.service.GetAllHelper(page, pageSize)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": helpers,
        "meta": gin.H{
            "page":       page,
            "page_size":  pageSize,
            "total":      total,
        },
    })
}