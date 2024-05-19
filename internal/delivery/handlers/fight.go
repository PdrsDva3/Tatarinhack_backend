package handlers

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FightHandler struct {
	service service.FightService
}

func InitFightHandler(service service.FightService) FightHandler {
	return FightHandler{
		service: service,
	}
}

// @Summary Get answer
// @Tags fight
// @Accept  json
// @Produce  json
// @Param id query int true "AnswerID"
// @Success 200 {object} int "Successfully get answer"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /fight/{id} [get]
func (p FightHandler) GetFight(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	teach, val, err := p.service.Get(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"session": id, "test": teach, "id_1": 100, "id_2": 101, "res_1": 5, "res_2": val})
}

// @Summary Create answer
// @Tags fight
// @Accept  json
// @Produce  json
// @Param data body entities.FightStart true "answer create"
// @Success 200 {object} int "Successfully created answer"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /fight/create [post]
func (p FightHandler) Post(c *gin.Context) {
	var answerCreate entities.FightStart

	if err := c.ShouldBindJSON(&answerCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err := p.service.Post(ctx, answerCreate.Pupa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": 100})
}
