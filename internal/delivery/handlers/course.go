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

type CourseHandler struct {
	service service.CourseService
}

func InitCourseHandler(service service.CourseService) CourseHandler {
	return CourseHandler{
		service: service,
	}
}

// @Summary Create course
// @Tags course
// @Accept  json
// @Produce  json
// @Param data body entities.CourseBase true "course create"
// @Success 200 {object} int "Successfully created course"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /course/create [post]
func (p CourseHandler) CreateCourse(c *gin.Context) {
	var create entities.CourseBase

	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := p.service.Create(ctx, create)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Add test to course
// @Tags course
// @Accept  json
// @Produce  json
// @Param data body entities.CourseAdd true "course add test"
// @Success 200 {object} int "Successfully add course"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /course/add [put]
func (p CourseHandler) AddCourse(c *gin.Context) {
	var change entities.CourseAdd

	if err := c.ShouldBindJSON(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err := p.service.AddTest(ctx, change.IDCourse, change.IDTest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"add": "success"})
}

// @Summary Get course
// @Tags course
// @Accept  json
// @Produce  json
// @Param id query int true "CourseID"
// @Success 200 {object} int "Successfully get course"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /course/{id} [get]
func (p CourseHandler) GetCourse(c *gin.Context) {
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
