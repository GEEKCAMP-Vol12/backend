package http

import (
	"net/http"
	"strconv"

	"github.com/taku10101/internal/domain"
	"github.com/taku10101/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SleepHandler struct {
    sleepUseCase *usecase.SleepUseCase
}
func NewSleepHandler(r *gin.Engine, sleepUseCase *usecase.SleepUseCase) {
    handler := &SleepHandler{sleepUseCase: sleepUseCase}
    r.GET("/sleeps/user/:id", handler.GetSleepsByUserID)
    r.GET("/sleeps/:id", handler.GetSleepByID)
    r.POST("/sleeps", handler.CreateSleep)
    r.PUT("/sleeps/:id", handler.UpdateSleep)
    r.DELETE("/sleeps/:id", handler.DeleteSleep)
    r.GET("/sleeps/user/:id/score", handler.GetSleepScore)
    r.GET("/sleeps/user/:id/last7days", handler.GetSleepsLast7Days)

}

func (h *SleepHandler) GetSleepsByUserID(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    sleeps, err := h.sleepUseCase.GetSleepsByUserID(strconv.Itoa(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sleeps)
}

func (h *SleepHandler) GetSleepByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sleep ID"})
        return
    }

    sleep, err := h.sleepUseCase.GetSleepByID(strconv.Itoa(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sleep)
}

func (h *SleepHandler) CreateSleep(c *gin.Context) {
    var sleep domain.Sleep
    if err := c.ShouldBindJSON(&sleep); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.sleepUseCase.CreateSleep(&sleep); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, sleep)
}

func (h *SleepHandler) UpdateSleep(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sleep ID"})
        return
    }

    var sleep domain.Sleep
    if err := c.ShouldBindJSON(&sleep); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    sleep.ID = id
    if err := h.sleepUseCase.UpdateSleep(&sleep); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sleep)
}

func (h *SleepHandler) DeleteSleep(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sleep ID"})
        return
    }

    if err := h.sleepUseCase.DeleteSleep(strconv.Itoa(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Sleep deleted"})
}

func (h *SleepHandler) GetSleepScore(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    score, err := h.sleepUseCase.GetSleepScore(strconv.Itoa(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"score": score})
}
func (h *SleepHandler) GetSleepsLast7Days(c *gin.Context) {
    userID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    sleeps, err := h.sleepUseCase.GetSleepsLast7Days(strconv.Itoa(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sleeps)
}

