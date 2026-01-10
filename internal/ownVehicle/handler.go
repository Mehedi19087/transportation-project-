package ownvehicle

import (
    "net/http"
    "strconv"

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

// ========== OwnVehicleTrip Handler ==========

type OwnVehicleTripHandler struct {
    service OwnVehicleTripService
}

func NewOwnVehicleTripHandler(service OwnVehicleTripService) *OwnVehicleTripHandler {
    return &OwnVehicleTripHandler{service: service}
}

// POST /api/v1/own-vehicle-trips
func (h *OwnVehicleTripHandler) CreateOwnVehicleTrip(ctx *gin.Context) {
    var req CreateOwnVehicleTripDTO

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.Create(&req); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "own vehicle trip created successfully"})
}

// GET /api/v1/own-vehicle-trips/:id
func (h *OwnVehicleTripHandler) GetOwnVehicleTrip(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    trip, err := h.service.Get(uint(id))
    if err != nil {
        if err.Error() == "own vehicle trip not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": trip})
}

// GET /api/v1/own-vehicle-trips?page=1&page_size=10
func (h *OwnVehicleTripHandler) GetAllOwnVehicleTrips(ctx *gin.Context) {
    pageStr := ctx.DefaultQuery("page", "1")
    sizeStr := ctx.DefaultQuery("page_size", "10")

    page, _ := strconv.Atoi(pageStr)
    pageSize, _ := strconv.Atoi(sizeStr)

    trips, total, err := h.service.GetAll(page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": trips,
        "meta": gin.H{
            "page":        page,
            "page_size":   pageSize,
            "total":       total,
            "total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
        },
    })
}

// PUT /api/v1/own-vehicle-trips/:id
func (h *OwnVehicleTripHandler) UpdateOwnVehicleTrip(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var req UpdateOwnVehicleTripDTO
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.Update(uint(id), &req); err != nil {
        if err.Error() == "own vehicle trip not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

// DELETE /api/v1/own-vehicle-trips/:id
func (h *OwnVehicleTripHandler) DeleteOwnVehicleTrip(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    if err := h.service.Delete(uint(id)); err != nil {
        if err.Error() == "own vehicle trip not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

// GET /api/v1/own-vehicle-trips/by-vehicle?vehicle_no=...&start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
func (h *OwnVehicleTripHandler) GetOwnVehicleTripsByVehicle(ctx *gin.Context) {
    vehicleNo := ctx.Query("vehicle_no")
    startDate := ctx.Query("start_date")
    endDate := ctx.Query("end_date")

    if vehicleNo == "" || startDate == "" || endDate == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "vehicle_no, start_date, end_date are required"})
        return
    }

    trips, err := h.service.GetByVehicleAndDateRange(vehicleNo, startDate, endDate)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": trips,
        "meta": gin.H{
            "vehicle_no": vehicleNo,
            "start_date": startDate,
            "end_date":   endDate,
            "count":      len(trips),
        },
    })
}