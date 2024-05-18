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

type QuestionHandler struct {
	service service.QuestionService
}

func InitQuestionHandler(service service.QuestionService) QuestionHandler {
	return QuestionHandler{
		service: service,
	}
}

// @Summary Create Que
// @Tags Que
// @Accept  json
// @Produce  json
// @Param data body entities.QuestionBase true "Que create"
// @Success 200 {object} int "Successfully created Que"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /question/create [post]
func (p QuestionHandler) CreateQue(c *gin.Context) {
	var QueCreate entities.QuestionBase

	if err := c.ShouldBindJSON(&QueCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := p.service.Create(ctx, QueCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Add Answer to Que
// @Tags Que
// @Accept  json
// @Produce  json
// @Param data body entities.QuestionAdd true "Que add ans"
// @Success 200 {object} int "Successfully add ans"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /question/add [put]
func (p QuestionHandler) AddAnswer(c *gin.Context) {
	var QueChange entities.QuestionAdd

	if err := c.ShouldBindJSON(&QueChange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err := p.service.AddAnswer(ctx, QueChange.IDQuestion, QueChange.IDAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"add": "success"})
}

// @Summary Get Que
// @Tags Que
// @Accept  json
// @Produce  json
// @Param id query int true "QueID"
// @Success 200 {object} int "Successfully get Que"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /question/{id} [get]
func (p QuestionHandler) GetQue(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"Que": teach})
}
