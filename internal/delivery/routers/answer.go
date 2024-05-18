package routers

import (
	"Tatarinhack_backend/internal/delivery/handlers"
	"Tatarinhack_backend/internal/repository/answer"
	answerserv "Tatarinhack_backend/internal/service/answer"
	"Tatarinhack_backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterAnswerRouter(r *gin.Engine, db *sqlx.DB, logs *logger.Logs) *gin.RouterGroup {

	answerRouter := r.Group("/answer")

	answerRepo := answer.InitAnswerRepository(db)
	answerService := answerserv.InitAnswerService(answerRepo)
	answerHandler := handlers.InitAnswerHandler(answerService)

	answerRouter.POST("/create", answerHandler.CreateAnswer)
	answerRouter.GET("/:id", answerHandler.GetAnswer)
	answerRouter.PUT("/change", answerHandler.ChangeAnswer)
	return answerRouter
}
