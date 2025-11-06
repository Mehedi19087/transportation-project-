package purchase

import (
    "net/http"
    "strconv"
    "transportation/internal/media"

    "github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
    service  PurchaseService
    uploader media.Uploader
}

func NewPurchaseHandler(service PurchaseService, uploader media.Uploader) *PurchaseHandler {
    return &PurchaseHandler{service: service, uploader: uploader}
}

func (h *PurchaseHandler) CreatePurchase(ctx *gin.Context) {
    var req CreatePurchase
    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }

    // optional file upload ("bill_image" like your field; also accept "image" for consistency)
    if fh, err := ctx.FormFile("bill_image"); err == nil && fh != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fh, "purchases")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "bill image upload failed"})
            return
        }
        req.BillImage = &url
    } else if fh, err := ctx.FormFile("image"); err == nil && fh != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fh, "purchases")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "bill image upload failed"})
            return
        }
        req.BillImage = &url
    }

    if err := h.service.CreatePurchase(&req); err != nil {
        ctx.JSON(http.StatusInternalServerError, err.Error())
        return
    }
    ctx.JSON(http.StatusCreated, "successfully created purchase")
}

func (h *PurchaseHandler) GetPurchase(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase ID"})
        return
    }

    res, err := h.service.GetPurchase(uint(id))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "purchase not found"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *PurchaseHandler) UpdatePurchase(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase ID"})
        return
    }

    var req UpdatePurchase
    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // optional bill image update
    if fh, err := ctx.FormFile("bill_image"); err == nil && fh != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fh, "purchases")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "bill image upload failed"})
            return
        }
        req.BillImage = &url
    } else if fh, err := ctx.FormFile("image"); err == nil && fh != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fh, "purchases")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "bill image upload failed"})
            return
        }
        req.BillImage = &url
    }

    if err := h.service.UpdatePurchase(uint(id), &req); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

func (h *PurchaseHandler) DeletePurchase(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil || id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase ID"})
        return
    }

    if err := h.service.DeletePurchase(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

func (h *PurchaseHandler) GetAllPurchase(ctx *gin.Context) {
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

    items, total, err := h.service.GetAllPurchase(page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{
        "data": items,
        "meta": gin.H{
            "page":      page,
            "page_size": pageSize,
            "total":     total,
        },
    })
}