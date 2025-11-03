package dealer

import (
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DealerHanlder struct {
	service DealerService
}

func NewDealerHandler(service DealerService) *DealerHanlder {
	return &DealerHanlder{service: service}
}

func (h *DealerHanlder) CreateDealer(ctx *gin.Context) {
	 var req DealerReq
	 if err := ctx.ShouldBindJSON(&req); err != nil {
		 ctx.JSON(400, gin.H{
			"message":err.Error(),
		 })
	 }
	 err:= h.service.CreateDealer(&req) 
	 if err != nil {
		 ctx.JSON(500, gin.H{
			 "message":err.Error(),
		 })
		 return 
	 }
	 ctx.JSON(201,"successfully created")
}


func (h *DealerHanlder) GetDealer(ctx *gin.Context) {
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
	res , err := h.service.GetDealer(uint(idStr))

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


func (h *DealerHanlder) UpdateDealer(ctx *gin.Context) {
     var req DealerUpdateReq
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
	if err := ctx.ShouldBindJSON(&req) ; err != nil {
		 ctx.JSON(400, err.Error())
		 return 
	}

	err = h.service.UpdateDealer(uint(idStr),&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	ctx.JSON(http.StatusAccepted,"successfully updated")

}

func(h *DealerHanlder) DeleteDealer(ctx *gin.Context) {
	 id:= ctx.Param("id")
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
	err =h.service.DeleteDealer(uint(idStr))
	if err != nil {
		 ctx.JSON(http.StatusInternalServerError, err.Error())
		 return 
	}
	ctx.JSON(http.StatusOK,"successfully deleted")
}

func(h *DealerHanlder) GetAllDealer(ctx *gin.Context) {
	pageStr:= ctx.DefaultQuery("page","1")
	pageSizeStr:= ctx.DefaultQuery("page_size","10")

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
	 dealers, total, err := h.service.GetAllDealer(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": dealers,
		"meta": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}