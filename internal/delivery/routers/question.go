package routers

import (
	"Tatarinhack_backend/internal/delivery/handlers"
	"Tatarinhack_backend/internal/repository/question"
	questionserv "Tatarinhack_backend/internal/service/question"
	"Tatarinhack_backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterQuestionRouter(r *gin.Engine, db *sqlx.DB, logs *logger.Logs) *gin.RouterGroup {

	questionRouter := r.Group("/question")

	questionRepo := question.InitQuestionRepository(db)
	questionService := questionserv.InitQuestService(questionRepo)
	questionHandler := handlers.InitQuestionHandler(questionService)

	questionRouter.POST("/create", questionHandler.CreateQue)
	questionRouter.GET("/:id", questionHandler.GetQue)
	questionRouter.PUT("/add", questionHandler.AddAnswer)
	return questionRouter
}
