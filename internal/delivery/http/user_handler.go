package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taku10101/internal/domain"
	"github.com/taku10101/internal/usecase"
)

type UserHandler struct {
    userUseCase *usecase.UserUseCase
}

func NewUserHandler(r *gin.Engine, userUseCase *usecase.UserUseCase) {
    handler := &UserHandler{userUseCase: userUseCase}
    r.GET("/users", handler.GetUsers)
    r.GET("/users/:id", handler.GetUserByID)
    r.POST("/users", handler.CreateUser)
    r.PUT("/users/:id", handler.UpdateUser)
    r.DELETE("/users/:id", handler.DeleteUser)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
    users, err := h.userUseCase.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }


    user, err := h.userUseCase.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req domain.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.userUseCase.CreateUser(&req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, req)
}


func (h *UserHandler) UpdateUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var req domain.UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    req.ID = strconv.Itoa(id)
    if err := h.userUseCase.UpdateUser(&req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, req)

}

func (h *UserHandler) DeleteUser(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := h.userUseCase.DeleteUser(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}