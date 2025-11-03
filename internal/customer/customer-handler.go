package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type CustomerHanlder struct {
	 service CustomerService
}

func NewCustomerHandler(service CustomerService) *CustomerHanlder {
	 return &CustomerHanlder{service: service}
}

func (h *CustomerHanlder) CreateCustomer(ctx *gin.Context) {
	 var req CreateCustomerRequest 
	 if err := ctx.ShouldBindJSON(&req) ; err != nil {
		  ctx.JSON(http.StatusBadRequest, gin.H{
			 "message":"please give customer name",
		  })
	 }
	 err := h.service.CreateCustomer(&req)
	 if err != nil {
		 ctx.JSON(500,err)
	 }
	 ctx.JSON(201,"customer created successfully")
}

func(h *CustomerHanlder) UpdateCustomer(ctx *gin.Context) {
	 id:=ctx.Param("id")
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
	 var req UpdateCustomerRequest 
	 if err:= ctx.ShouldBindJSON(&req) ; err != nil {
		  ctx.JSON(http.StatusBadRequest, gin.H{
			  "message" :"please give me value for customer",
		  })
		  return 
	 }
	 err = h.service.UpdateCustomer(&req, uint(idStr)) 
	 if err != nil {
		 ctx.JSON(500,err)
	 }
	 ctx.JSON(201,"customer created successfully")

}

func(h *CustomerHanlder) GetCustomer(ctx *gin.Context) {
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
	res , err := h.service.GetCustomer(uint(idStr))

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

func (h *CustomerHanlder) GetAllCustomer(ctx *gin.Context) {
	 pageStr:= ctx.DefaultQuery("page", "1")
	 pageSizeStr:=ctx.DefaultQuery("page_size","20")
	 
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
	
    customers, total, err := h.service.ListCustomers(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": customers,
		"meta": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})
}

func (h *CustomerHanlder) DeleteCustomer(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid customer id",
		})
		return
	}

	if err := h.service.DeleteCustomer(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}
