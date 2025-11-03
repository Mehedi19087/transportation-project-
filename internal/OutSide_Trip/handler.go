package outsidetrip

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type OutSideTripHandler struct {
    service OutSideTripService
}

func NewOutSideTripHandler(service OutSideTripService) *OutSideTripHandler {
    return &OutSideTripHandler{service: service}
}

func (h *OutSideTripHandler) CreateOutSideTrip(ctx *gin.Context) {
    var req OutSideTripReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(400, gin.H{
            "message": err.Error(),
        })
    }
    err := h.service.CreateOutSideTrip(&req)
    if err != nil {
        ctx.JSON(500, gin.H{
            "message": err.Error(),
        })
        return
    }
    ctx.JSON(201, "successfully created")
}

func (h *OutSideTripHandler) GetOutSideTrip(ctx *gin.Context) {
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid note ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id can not become zero",
        })
        return
    }
    res, err := h.service.GetOutSideTrip(uint(idStr))

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "server error",
        })
        return
    }

    ctx.JSON(http.StatusAccepted, gin.H{
        "note": res,
    })
}

func (h *OutSideTripHandler) UpdateOutSideTrip(ctx *gin.Context) {
    var req OutSideTripUpdateReq
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid note ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id can not become zero",
        })
        return
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(400, err.Error())
        return
    }

    err = h.service.UpdateOutSideTrip(uint(idStr), &req)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "server error",
        })
        return
    }

    ctx.JSON(http.StatusAccepted, "successfully updated")
}

func (h *OutSideTripHandler) DeleteOutSideTrip(ctx *gin.Context) {
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid note ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id can not become zero",
        })
        return
    }
    err = h.service.DeleteOutSideTrip(uint(idStr))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err.Error())
        return
    }
    ctx.JSON(http.StatusOK, "successfully deleted")
}

func (h *OutSideTripHandler) GetAllOutSideTrip(ctx *gin.Context) {
    pageStr := ctx.DefaultQuery("page", "1")
    pageSizeStr := ctx.DefaultQuery("page_size", "10")

    page, err := strconv.Atoi(pageStr)

    if err != nil || page < 1 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "page must be a positive integer",
        })
        return
    }
    pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil || pageSize < 1 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "page_size must be a positive integer",
        })
        return
    }
    trips, total, err := h.service.GetAllOutSideTrip(page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "server error",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data":  trips,
        "meta": gin.H{
            "page":      page,
            "page_size": pageSize,
            "total":     total,
        },
    })
}

func (h *OutSideTripHandler) GetVehicleMonthlySummary(ctx *gin.Context) {
    vehicleNumber := ctx.Param("vehicleNumber")
    month := ctx.Param("month")
    
    if vehicleNumber == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "vehicle number is required"})
        return
    }
    if month == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "month is required"})
        return
    }
    
    summary, err := h.service.GetVehicleMonthlySummary(vehicleNumber, month)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"data": summary})
}