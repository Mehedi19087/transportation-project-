package calan

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CalanHandler struct {
	service CalanService
}

func NewCalanHandler(service CalanService) *CalanHandler {
	return &CalanHandler{service: service}
}

func (h *CalanHandler) GetCalans(ctx *gin.Context) {
	var filters FilterCalanReq

    if err := ctx.ShouldBindQuery(&filters); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
        return
    }

    // Set defaults if not provided
    if filters.Page < 1 {
        filters.Page = 1
    }
    if filters.PageSize < 1 || filters.PageSize >10 {
        filters.PageSize = 10
    }

	calans, total, err := h.service.GetCalans(filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch calans"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": calans,
		"meta": gin.H{
			"page":      filters.Page,
			"page_size": filters.PageSize,
			"total":     total,
		},
	})
}

func (h *CalanHandler) UpdateStatus(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var req struct {
        Status string `json:"status" binding:"required"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.UpdateCalanStatus(uint(id), req.Status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}
