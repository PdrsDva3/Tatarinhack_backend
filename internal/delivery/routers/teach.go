package routers

import (
	"Tatarinhack_backend/internal/delivery/handlers"
	_ "Tatarinhack_backend/internal/delivery/middleware"
	"Tatarinhack_backend/internal/repository/teach"
	teachserv "Tatarinhack_backend/internal/service/teach"
	"Tatarinhack_backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterTeachRouter(r *gin.Engine, db *sqlx.DB, logs *logger.Logs) *gin.RouterGroup {

	teachRouter := r.Group("/teach")

	teachRepo := teach.InitTeachRepository(db)
	teachService := teachserv.InitTeachService(teachRepo)
	teachHandler := handlers.InitTeachHandler(teachService)

	teachRouter.POST("/create", teachHandler.CreateAdmin)
	teachRouter.POST("/login", teachHandler.LoginTeach)
	teachRouter.GET("/:id", teachHandler.GetTeach)
	teachRouter.PUT("/change", teachHandler.ChangePWD)
	teachRouter.DELETE("/delete/:id", teachHandler.DeleteTeach)
	return teachRouter
}
