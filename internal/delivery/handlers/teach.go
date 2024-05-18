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

type TeachHandler struct {
	service service.TeachService
}

func InitTeachHandler(service service.TeachService) TeachHandler {
	return TeachHandler{
		service: service,
	}
}

// @Summary Create teach
// @Tags teach
// @Accept  json
// @Produce  json
// @Param data body entities.TeachCreate true "teach create"
// @Success 200 {object} int "Successfully created teach"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /teach/create [post]
func (p TeachHandler) CreateAdmin(c *gin.Context) {
	var teachCreate entities.TeachCreate

	if err := c.ShouldBindJSON(&teachCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := p.service.Create(ctx, teachCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Login teach
// @Tags teach
// @Accept  json
// @Produce  json
// @Param data body entities.TeachLogin true "teach login"
// @Success 200 {object} int "Successfully login teach"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /teach/login [post]
func (p TeachHandler) LoginTeach(c *gin.Context) {
	var teachLogin entities.TeachLogin

	if err := c.ShouldBindJSON(&teachLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := p.service.Login(ctx, teachLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get teach
// @Tags teach
// @Accept  json
// @Produce  json
// @Param id query int true "TeachID"
// @Success 200 {object} int "Successfully get teach"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /teach/{id} [get]
func (p TeachHandler) GetTeach(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"teach": teach})
}

// @Summary ChangePWD teach
// @Tags teach
// @Accept  json
// @Produce  json
// @Param data body entities.TeachChangePassword true "teach change pwd"
// @Success 200 {object} int "Successfully change pwd"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /teach/change [put]
func (p TeachHandler) ChangePWD(c *gin.Context) {
	var changePWD entities.TeachChangePassword

	if err := c.ShouldBindJSON(&changePWD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err := p.service.ChangePassword(ctx, changePWD.ID, changePWD.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"change": "access"})
}

// @Summary Delete teach
// @Tags teach
// @Accept  json
// @Produce  json
// @Param id query int true "TeachID"
// @Success 200 {object} int "Successfully delete teach"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /teach/delete/{id} [delete]
func (p TeachHandler) DeleteTeach(c *gin.Context) {
	id := c.Query("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	err = p.service.Delete(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete": id})
}
