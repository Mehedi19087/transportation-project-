package ownvehicle

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Handler struct {
    service Service
}

func NewHandler(service Service) *Handler {
    return &Handler{service: service}
}

// GET /api/v1/own-vehicles/by-driver?driver_name=...&start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
func (h *Handler) GetByDriverAndDate(ctx *gin.Context) {
    var q DriverDateQuery
    if err := ctx.ShouldBindQuery(&q); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "driver_name, start_date, end_date are required"})
        return
    }
    items, err := h.service.GetOwnVehiclesByDriverAndDate(q.DriverName, q.StartDate, q.EndDate)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "data": items,
        "meta": gin.H{
            "driver_name": q.DriverName,
            "start_date":  q.StartDate,
            "end_date":    q.EndDate,
            "count":       len(items),
        },
    })
}