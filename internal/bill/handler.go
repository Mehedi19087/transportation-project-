package bill

import (
	"net/http"
	"strconv"
	"errors"
	"gorm.io/gorm"
    // "encoding/json"
    // "log"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context){
	 var req CreateProductReq
	 if err := ctx.ShouldBindJSON(&req) ; err != nil {
		 ctx.JSON(http.StatusBadRequest,"name is required")
		 return 
	 }
	 err := h.service.CreateProduct(&req) 
	 if err != nil {
		 ctx.JSON(http.StatusInternalServerError,err.Error())
		 return 
	 }
	 ctx.JSON(201, "successfully created")
}
func(h *ProductHandler) GetProduct(ctx *gin.Context) {
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
	res , err := h.service.GetProduct(uint(idStr))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"note": res,
	})
}

func(h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	 id:= ctx.Param("id")
	 idStr, err := strconv.ParseUint(id, 10, 32)
	 if err != nil {
		 ctx.JSON(http.StatusBadRequest, err.Error())
		 return 
	 }
	 if idStr == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id can not become zero",
		})
		return
	}
	var req UpdateProductReq 
    if err := ctx.ShouldBindJSON(&req); err!= nil {
		ctx.JSON(http.StatusBadRequest,err.Error())
		return 
	}
	err = h.service.UpdateProduct(uint(idStr),&req)
	if err != nil {
		 ctx.JSON(http.StatusInternalServerError, err.Error())
		 return
	}
	ctx.JSON(200, "successfully updated")
}

func(h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	 id:= ctx.Param("id")
	 idStr, err := strconv.ParseUint(id, 10, 32)
	 if err != nil {
		 ctx.JSON(http.StatusBadRequest, err.Error())
		 return
	 }
	 if idStr ==0 {
		 ctx.JSON(http.StatusBadRequest, errors.New("select specific product for delete")) 
	 }
	 err =h.service.DeleteProduct(uint(idStr))
	 if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	 }
	 ctx.JSON(http.StatusOK, "successfully deleted")

}

func(h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	pageStr:= ctx.DefaultQuery("page", "1")
	pageSizeStr:= ctx.DefaultQuery("pageSize", "10")
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
	 products, total, err := h.service.GetAllProducts(page, pageSize)
	 if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
		"meta": gin.H{
			"page":      page,
			"page_size": pageSize,
			"total":     total,
		},
	})


}

func (h *ProductHandler) GetCompanyProducts(ctx *gin.Context) {
    companyIDParam := ctx.Param("companyId")
    cid, err := strconv.ParseUint(companyIDParam, 10, 32)
    if err != nil || cid == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid company id"})
        return
    }
    
    pageStr := ctx.DefaultQuery("page", "1")
    sizeStr := ctx.DefaultQuery("page_size", "10")
    page, _ := strconv.Atoi(pageStr)
    size, _ := strconv.Atoi(sizeStr)
    
    items, total, err := h.service.GetCompanyProducts(uint(cid), page, size)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{
        "data": items,
        "meta": gin.H{"page": page, "page_size": size, "total": total},
    })
}

// ...existing code...

func (h *ProductHandler) CreateBill(ctx *gin.Context) {
    var req CreateBillReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    // if b, err := json.MarshalIndent(req, "", "  "); err == nil {
    //     log.Println("CreateBill payload:\n" + string(b))
    // }

    err := h.service.CreateBill(&req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "bill created successfully",
    })
}

// ...existing code...

func (h *ProductHandler) GetBill(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID format",
        })
        return
    }

    bill, err := h.service.GetBill(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            ctx.JSON(http.StatusNotFound, gin.H{
                "error": "Bill not found",
            })
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": bill,
    })
}

// ...existing code...

func (h *ProductHandler) GetProductBills(ctx *gin.Context) {
    productIDParam := ctx.Param("id")
    productID, err := strconv.ParseUint(productIDParam, 10, 32)
    if err != nil || productID == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid product id",
        })
        return
    }

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

    bills, total, err := h.service.GetProductBills(uint(productID), page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": bills,
        "meta": gin.H{
            "page":      page,
            "page_size": pageSize,
            "total":     total,
        },
    })
}

// ...existing code...

func (h *ProductHandler) UpdateBill(ctx *gin.Context) {
    var req UpdateBillReq
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID format",
        })
        return
    }
    if id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    err = h.service.UpdateBill(uint(id), &req)
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

// ...existing code...

func (h *ProductHandler) DeleteBill(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid ID format",
        })
        return
    }
    if id == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "id cannot be zero",
        })
        return
    }

    err = h.service.DeleteBill(uint(id))
    if err != nil {
        if err.Error() == "bill not found" {
            ctx.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "successfully deleted",
    })
}

// ...existing code...

func (h *ProductHandler) UpdateProductTripFields(ctx *gin.Context) {
    var req UpdateProductTripFieldsReq
    productIDParam := ctx.Param("id")
    productID, err := strconv.ParseUint(productIDParam, 10, 32)
    if err != nil || productID == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid product id",
        })
        return
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    err = h.service.UpdateProductTripFields(uint(productID), &req)
    if err != nil {
        if err.Error() == "product not found" {
            ctx.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "trip fields updated successfully",
    })
}

// ...existing code...

func (h *ProductHandler) GetProductTripFields(ctx *gin.Context) {
    productIDParam := ctx.Param("id")
    productID, err := strconv.ParseUint(productIDParam, 10, 32)
    if err != nil || productID == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid product id",
        })
        return
    }

    tripFields, err := h.service.GetProductTripFields(uint(productID))
    if err != nil {
        if err.Error() == "product not found" {
            ctx.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": tripFields,
    })
}

// ...existing code...

func (h *ProductHandler) UpdateProductBillFields(ctx *gin.Context) {
    var req UpdateProductBillFieldsReq
    productIDParam := ctx.Param("id")
    productID, err := strconv.ParseUint(productIDParam, 10, 32)
    if err != nil || productID == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid product id",
        })
        return
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    err = h.service.UpdateProductBillFields(uint(productID), &req)
    if err != nil {
        if err.Error() == "product not found" {
            ctx.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "bill fields updated successfully",
    })
}

func (h *ProductHandler) GetProductBillFields(ctx *gin.Context) {
    productIDParam := ctx.Param("id")
    productID, err := strconv.ParseUint(productIDParam, 10, 32)
    if err != nil || productID == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid product id",
        })
        return
    }

    billFields, err := h.service.GetProductBillFields(uint(productID))
    if err != nil {
        if err.Error() == "product not found" {
            ctx.JSON(http.StatusNotFound, gin.H{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": billFields,
    })
}