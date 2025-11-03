package trip

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Handler struct {
    service Service
}

func NewHandler(s Service) *Handler {
    return &Handler{service: s}
}

func (h *Handler) CreateTrip(ctx *gin.Context) {
    var req CreateTripReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.CreateTrip(&req); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": "trip created successfully"})
}

func (h *Handler) GetTrip(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    res, err := h.service.GetTrip(uint(id))
    if err != nil {
        if err.Error() == "trip not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) UpdateTrip(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    var req UpdateTripReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.UpdateTrip(uint(id), &req); err != nil {
        if err.Error() == "trip not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

func (h *Handler) DeleteTrip(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    if err := h.service.DeleteTrip(uint(id)); err != nil {
        if err.Error() == "trip not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

func (h *Handler) GetProductTrips(ctx *gin.Context) {
    productIDParam := ctx.Param("id")
    pid, err := strconv.ParseUint(productIDParam, 10, 32)
    if err != nil || pid == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
        return
    }
    pageStr := ctx.DefaultQuery("page", "1")
    sizeStr := ctx.DefaultQuery("page_size", "10")
    page, _ := strconv.Atoi(pageStr)
    size, _ := strconv.Atoi(sizeStr)

    items, total, err := h.service.GetProductTrips(uint(pid), page, size)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "data": items,
        "meta": gin.H{"page": page, "page_size": size, "total": total},
    })
}