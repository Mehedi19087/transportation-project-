package vehicle

import (
	"net/http"
	"strconv"
	"transportation/internal/media"

	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
    service VehicleService
    uploader media.Uploader
}

func NewVehicleHandler(service VehicleService, uploader media.Uploader) *VehicleHandler {
    return &VehicleHandler{service: service, uploader: uploader}
}

func (h *VehicleHandler) CreateVehicle(ctx *gin.Context) {
    var req CreateVehicle
    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	if fileHeader, err := ctx.FormFile("image"); err == nil && fileHeader != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fileHeader, "vehicles")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "image upload failed"})
            return
        }
        req.ImageURL = &url
    }

    err := h.service.CreateVehicle(&req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": "successfully created vehicle"})
}

func (h *VehicleHandler) GetVehicle(ctx *gin.Context) {
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "id cannot be zero"})
        return
    }

    res, err := h.service.GetVehicle(uint(idStr))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "vehicle not found"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *VehicleHandler) UpdateVehicle(ctx *gin.Context) {
    var req UpdateVehicle
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "id cannot be zero"})
        return
    }

    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if fileHeader, err := ctx.FormFile("image"); err == nil && fileHeader != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fileHeader, "vehicles")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "image upload failed"})
            return
        }
        req.ImageURL = &url
    }

    err = h.service.UpdateVehicle(uint(idStr), &req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

func (h *VehicleHandler) DeleteVehicle(ctx *gin.Context) {
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vehicle ID"})
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "id cannot be zero"})
        return
    }

    err = h.service.DeleteVehicle(uint(idStr))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

func (h *VehicleHandler) GetAllVehicle(ctx *gin.Context) {
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

    vehicles, total, err := h.service.GetAllVehicle(page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": vehicles,
        "meta": gin.H{
            "page":      page,
            "page_size": pageSize,
            "total":     total,
        },
    })
}