package auth

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
    return &AuthHandler{service: service}
}

// POST /auth/register
func (h *AuthHandler) CreateUser(ctx *gin.Context) {
    var req AuthReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "name and password are required",
        })
        return
    }
    
    if err := h.service.CreateUser(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    ctx.JSON(http.StatusCreated, gin.H{
        "message": "user created successfully, waiting for admin approval",
    })
}

// POST /auth/login
func (h *AuthHandler) Login(ctx *gin.Context) {
    var req AuthReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "name and password are required",
        })
        return
    }
    
    token, role, err := h.service.Login(&req)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{
        "token": token,
        "role":  role,
        "message": "login successful",
    })
}

// PUT /auth/users/:id (for admin to approve or manager to update own profile)
func (h *AuthHandler) UpdateUser(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }
    
    var req UpdateReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }
    
    // Get requester role from JWT middleware (you'll add this)
    requesterRole, _ := ctx.Get("role")
    role := requesterRole.(string)
    
    if err := h.service.UpdateUser(uint(id), req, role); err != nil {
        ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}


func (h *AuthHandler) GetPendingUsers(ctx *gin.Context) {
    users, err := h.service.GetPendingUsers()
       if err != nil {
           ctx.JSON(http.StatusInternalServerError,gin.H {
			"error": err.Error(),
		   })
		   return 
		}
       ctx.JSON(http.StatusOK, users)
}