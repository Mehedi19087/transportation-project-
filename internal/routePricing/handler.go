package routepricing

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type RoutePricingHandler struct {
    service RoutePricingService
}

func NewRoutePricingHandler(service RoutePricingService) *RoutePricingHandler {
    return &RoutePricingHandler{service: service}
}

func (h *RoutePricingHandler) CreateRoutePricing(ctx *gin.Context) {
    var req RoutePricingReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    err := h.service.CreateRoutePricing(&req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": "successfully created"})
}

func (h *RoutePricingHandler) GetRoutePricing(ctx *gin.Context) {
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid route pricing ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }
    res, err := h.service.GetRoutePricing(uint(idStr))

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "route pricing not found",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": res,
    })
}

func (h *RoutePricingHandler) UpdateRoutePricing(ctx *gin.Context) {
    var req RoutePricingUpdateReq
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid route pricing ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = h.service.UpdateRoutePricing(uint(idStr), &req)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

func (h *RoutePricingHandler) DeleteRoutePricing(ctx *gin.Context) {
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid route pricing ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }

    err = h.service.DeleteRoutePricing(uint(idStr))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

func (h *RoutePricingHandler) GetAllRoutePricing(ctx *gin.Context) {
    pageStr := ctx.DefaultQuery("page", "1")
    pageSizeStr := ctx.DefaultQuery("page_size", "10")

    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }

    pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil || pageSize < 1 {
        pageSize = 10
    }

    routePricings, total, err := h.service.GetAllRoutePricing(page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": routePricings,
        "meta": gin.H{
            "page":      page,
            "page_size": pageSize,
            "total":     total,
        },
    })
}

func (h *RoutePricingHandler) GetRate(c *gin.Context) {
    var req RateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    rate, err := h.service.GetRateByLocations(req.LoadPoint, req.UnloadPoint)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"rate": rate})
}