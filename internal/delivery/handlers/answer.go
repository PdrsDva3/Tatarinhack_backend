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

type AnswerHandler struct {
	service service.AnswerService
}

func InitAnswerHandler(service service.AnswerService) AnswerHandler {
	return AnswerHandler{
		service: service,
	}
}

// @Summary Create answer
// @Tags answer
// @Accept  json
// @Produce  json
// @Param data body entities.AnswerBase true "answer create"
// @Success 200 {object} int "Successfully created answer"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /answer/create [post]
func (p AnswerHandler) CreateAnswer(c *gin.Context) {
	var answerCreate entities.AnswerBase

	if err := c.ShouldBindJSON(&answerCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := p.service.Create(ctx, answerCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Change iscorrect answer
// @Tags answer
// @Accept  json
// @Produce  json
// @Param data body entities.AnswerChange true "answer change"
// @Success 200 {object} int "Successfully change answer"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /answer/change [put]
func (p AnswerHandler) ChangeAnswer(c *gin.Context) {
	var answerChange entities.AnswerChange

	if err := c.ShouldBindJSON(&answerChange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err := p.service.Change(ctx, answerChange.ID, answerChange.IsCorrect)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"change": "success"})
}

// @Summary Get answer
// @Tags answer
// @Accept  json
// @Produce  json
// @Param id query int true "AnswerID"
// @Success 200 {object} int "Successfully get answer"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /answer/{id} [get]
func (p AnswerHandler) GetAnswer(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	teach, err := p.service.GetMe(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"answer": teach})
}
