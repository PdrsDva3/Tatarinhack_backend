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

type TestHandler struct {
	service service.TestService
}

func InitTestHandler(service service.TestService) TestHandler {
	return TestHandler{
		service: service,
	}
}

// @Summary Create test
// @Tags test
// @Accept  json
// @Produce  json
// @Param data body entities.TestBase true "test create"
// @Success 200 {object} int "Successfully created test"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /test/create [post]
func (p TestHandler) CreateTest(c *gin.Context) {
	var QueCreate entities.TestBase

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

// @Summary Add que to test
// @Tags test
// @Accept  json
// @Produce  json
// @Param data body entities.TestAdd true "test add que"
// @Success 200 {object} int "Successfully add test"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /test/add [put]
func (p TestHandler) AddTest(c *gin.Context) {
	var QueChange entities.TestAdd

	if err := c.ShouldBindJSON(&QueChange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err := p.service.AddTest(ctx, QueChange.IDQuestion, QueChange.IDTest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"add": "success"})
}

// @Summary Get test
// @Tags test
// @Accept  json
// @Produce  json
// @Param id query int true "TestID"
// @Success 200 {object} int "Successfully get test"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /test/{id} [get]
func (p TestHandler) GetTest(c *gin.Context) {
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

// @Summary put answer test
// @Tags test
// @Accept  json
// @Produce  json
// @Param data body entities.TestAnswer true "test answer"
// @Success 200 {object} int "Successfully test"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /test/answer [post]
func (p TestHandler) TestAnswer(c *gin.Context) {
	var QueCreate entities.TestAnswer

	if err := c.ShouldBindJSON(&QueCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	cnt, err := p.service.TestAnswer(ctx, QueCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"read": cnt})
}
