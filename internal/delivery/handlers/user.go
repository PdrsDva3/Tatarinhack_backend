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

type HandlerUser struct {
	service service.UserService
}

func InitUserHandler(service service.UserService) HandlerUser {
	return HandlerUser{
		service: service,
	}
}

// @Summary Create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body entities.UserCreate true "user create"
// @Success 200 {object} int "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/create [post]
func (handler HandlerUser) Create(g *gin.Context) {
	var newUser entities.UserCreate

	if err := g.ShouldBindJSON(&newUser); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := handler.service.Create(ctx, newUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Login user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body entities.UserLogin true "user login"
// @Success 200 {object} int "Successfully login user"
// @Failure 400 {object} map[string]string "Invalid password"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/login [post]
func (handler HandlerUser) Login(g *gin.Context) {
	var User entities.UserLogin

	if err := g.ShouldBindJSON(&User); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := handler.service.Login(ctx, User)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get user"
// @Failure 400 {object} map[string]string "Invalid UserID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/get/{id} [get]
func (handler HandlerUser) GetUser(g *gin.Context) {
	userID := g.Query("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := handler.service.GetUser(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary Get friend
// @Tags friend
// @Accept  json
// @Produce  json
// @Param id query int true "FriendID"
// @Success 200 {object} int "Successfully get friend"
// @Failure 400 {object} map[string]string "Invalid friendID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/get/friend/{id} [get]
func (handler HandlerUser) GetFriend(g *gin.Context) {
	friendID := g.Query("id")

	id, err := strconv.Atoi(friendID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := handler.service.GetFriend(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary Get man
// @Tags man
// @Accept  json
// @Produce  json
// @Param id query int true "ManID"
// @Success 200 {object} int "Successfully get man"
// @Failure 400 {object} map[string]string "Invalid manID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/get/man/{id} [get]
func (handler HandlerUser) GetMan(g *gin.Context) {
	manID := g.Query("id")

	id, err := strconv.Atoi(manID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := handler.service.GetMan(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary Adding friend
// @Tags friend
// @Accept  json
// @Produce  json
// @Param data body entities.UserAddFriend true "friend added"
// @Success 200 {object} int "Successfully add friend"
// @Failure 400 {object} map[string]string "Invalid IDs"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/add [post]
func (handler HandlerUser) AddFriend(g *gin.Context) {
	var adding entities.UserAddFriend

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := handler.service.AddFriend(ctx, adding.UserID, adding.FriendID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"ass": "success"})
}

// @Summary Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/delete/{id} [delete]
func (handler HandlerUser) Delete(g *gin.Context) {
	userID := g.Query("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = handler.service.Delete(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"delete": id})
}

// @Summary Change password
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body entities.UserChangePassword true "change password"
// @Success 200 {object} int "Success changing"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/pwd [put]
func (handler HandlerUser) UpdatePassword(g *gin.Context) {
	var user entities.UserChangePassword
	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := handler.service.UpdatePassword(ctx, user.ID, user.Password)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"change": "success"})
}

// @Summary Update grammar
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Success"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/gram/{id} [put]
func (handler HandlerUser) GrammarUp(g *gin.Context) {
	id, err := strconv.Atoi(g.Query("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = handler.service.GrammarUp(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"update": "success"})
}

// @Summary Update vocabulary
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Success"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/voc/{id} [put]
func (handler HandlerUser) VocabularyUp(g *gin.Context) {
	id, err := strconv.Atoi(g.Query("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = handler.service.VocabularyUp(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"update": "success"})
}

// @Summary Update speaking
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Success"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/speak/{id} [put]
func (handler HandlerUser) SpeakingUp(g *gin.Context) {
	id, err := strconv.Atoi(g.Query("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = handler.service.SpeakingUp(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"update": "success"})
}

// @Summary Get FriendsList
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get FriendsList"
// @Failure 400 {object} map[string]string "Invalid UserID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/get/friend/lst/{id} [get]
func (handler HandlerUser) FriendsList(g *gin.Context) {
	id, err := strconv.Atoi(g.Query("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	list, err := handler.service.GetFriendsList(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"list": list})
}

// @Summary Update level
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Success"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/lvl/{id} [put]
func (handler HandlerUser) LevelUp(g *gin.Context) {
	id, err := strconv.Atoi(g.Query("id"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = handler.service.SpeakingUp(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"update": "success"})
}
