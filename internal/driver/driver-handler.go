package driver

import (
	"net/http"
	"strconv"
	"transportation/internal/media"

	"github.com/gin-gonic/gin"
)


type DriverHandler struct {
	 service  DriverService
	 uploader media.Uploader
}

func NewDriverHandler(service DriverService, uploader media.Uploader) *DriverHandler {
    return &DriverHandler{service: service, uploader: uploader}
}

func (h *DriverHandler) CreateDriver(ctx *gin.Context) {
	 var req CreateDriver
	 if err := ctx.ShouldBind(&req); err != nil {
		 ctx.JSON(400, err.Error())
		 return 
	 }
     if fileHeader, err := ctx.FormFile("image"); err == nil && fileHeader != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fileHeader, "drivers")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "image upload failed"})
            return
        }
        req.ImageURL = &url
    }
	 err := h.service.CreateDriver(&req) 
	 if err != nil {
		 ctx.JSON(500,err.Error())
		 return 
	 }
	 ctx.JSON(201, "successfully created driver")
}

func (h *DriverHandler) GetDriver(ctx *gin.Context) {
	 id:= ctx.Param("id")
	 idStr, err := strconv.ParseUint(id, 10, 32)
	 if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid driver ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }
    
    res, err := h.service.GetDriver(uint(idStr))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "driver not found",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": res,
    })
}

func (h *DriverHandler) UpdateDriver(ctx *gin.Context) {
    var req UpdateDriver
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid driver ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }
    
    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
		}
    if fileHeader, err := ctx.FormFile("image"); err == nil && fileHeader != nil {
        url, err := h.uploader.UploadFromFileHeader(ctx, fileHeader, "drivers")
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "image upload failed"})
            return
        }
        req.ImageURL = &url
    }

    err = h.service.UpdateDriver(uint(idStr), &req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "successfully updated",
    })
}
func (h *DriverHandler) DeleteDriver(ctx *gin.Context) {
    id := ctx.Param("id")
    idStr, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid driver ID",
        })
        return
    }
    if idStr == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }
    
    err = h.service.DeleteDriver(uint(idStr))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
	ctx.JSON(http.StatusOK, gin.H{
        "message": "successfully deleted",
	})
}

func (h *DriverHandler) GetAllDriver(ctx *gin.Context) {
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

    drivers, total, err := h.service.GetAllDriver(page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
	 ctx.JSON(http.StatusOK, gin.H{
        "data": drivers,
        "meta": gin.H{
            "page":      page,
            "page_size": pageSize,
            "total":     total,
        },
    })
}