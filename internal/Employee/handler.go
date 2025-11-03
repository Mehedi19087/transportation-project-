package employee

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
    service EmployeeService
}

func NewEmployeeHandler(service EmployeeService) *EmployeeHandler {
    return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) CreateEmployee(ctx *gin.Context) {
    var req CreateEmployeeReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(400, gin.H{
            "message": err.Error(),
        })
        return
    }
    err := h.service.CreateEmployee(&req)
    if err != nil {
        ctx.JSON(500, gin.H{
            "message": err.Error(),
        })
        return
    }
    ctx.JSON(201, "successfully created")
}

func (h *EmployeeHandler) GetEmployee(ctx *gin.Context) {
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
    res, err := h.service.GetEmployee(uint(idStr))

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

func (h *EmployeeHandler) UpdateEmployee(ctx *gin.Context) {
    var req UpdateEmployeeReq
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

    err = h.service.UpdateEmployee(uint(idStr), &req)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "server error",
        })
        return
    }

    ctx.JSON(http.StatusAccepted, "successfully updated")
}

func (h *EmployeeHandler) DeleteEmployee(ctx *gin.Context) {
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
    err = h.service.DeleteEmployee(uint(idStr))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err.Error())
        return
    }
    ctx.JSON(http.StatusOK, "successfully deleted")
}

func (h *EmployeeHandler) GetAllEmployee(ctx *gin.Context) {
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
    employees, total, err := h.service.GetAllEmployee(page, pageSize)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "server error",
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": employees,
        "meta": gin.H{
            "page":      page,
            "page_size": pageSize,
            "total":     total,
        },
    })
}