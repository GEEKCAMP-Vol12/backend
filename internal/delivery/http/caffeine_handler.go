package http

import (
	"net/http"
	"strconv"

	"github.com/taku10101/internal/domain"
	"github.com/taku10101/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CaffeineHandler struct {
    caffeineUseCase *usecase.CaffeineUseCase
}

func NewCaffeineHandler(r *gin.Engine, caffeineUseCase *usecase.CaffeineUseCase) {
    handler := &CaffeineHandler{caffeineUseCase: caffeineUseCase}
    r.GET("/caffeines/user/:id", handler.GetCaffeinesByUserID)
    r.GET("/caffeines/:id", handler.GetCaffeineByID)
    r.POST("/caffeines", handler.CreateCaffeine)
    r.PUT("/caffeines/:id", handler.UpdateCaffeine)
    r.DELETE("/caffeines/:id", handler.DeleteCaffeine)
    r.GET("/caffeines/user/:id/score", handler.GetCaffeineScore)
}

func (h *CaffeineHandler) GetCaffeinesByUserID(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    caffeines, err := h.caffeineUseCase.GetCaffeinesByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, caffeines)
}

func (h *CaffeineHandler) GetCaffeineByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid caffeine ID"})
        return
    }

    caffeine, err := h.caffeineUseCase.GetCaffeineByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, caffeine)
}

func (h *CaffeineHandler) CreateCaffeine(c *gin.Context) {
    var caffeine domain.Caffeine
    if err := c.ShouldBindJSON(&caffeine); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.caffeineUseCase.CreateCaffeine(&caffeine); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, caffeine)
}

func (h *CaffeineHandler) UpdateCaffeine(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid caffeine ID"})
        return
    }

    var caffeine domain.Caffeine
    if err := c.ShouldBindJSON(&caffeine); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    caffeine.ID = id
    if err := h.caffeineUseCase.UpdateCaffeine(&caffeine); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, caffeine)
}

func (h *CaffeineHandler) DeleteCaffeine(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid caffeine ID"})
        return
    }

    if err := h.caffeineUseCase.DeleteCaffeine(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Caffeine deleted"})
}

func (h *CaffeineHandler) GetCaffeineScore(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    score, err := h.caffeineUseCase.GetCaffeineScore(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"score": score})
}
