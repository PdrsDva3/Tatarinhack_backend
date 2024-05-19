package routers

import (
	"Tatarinhack_backend/internal/delivery/handlers"
	"Tatarinhack_backend/internal/repository/test"
	testserv "Tatarinhack_backend/internal/service/test"
	"Tatarinhack_backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterTestRouter(r *gin.Engine, db *sqlx.DB, logs *logger.Logs) *gin.RouterGroup {

	TestRouter := r.Group("/test")

	TestRepo := test.InitTestRepository(db)
	TestService := testserv.InitTestService(TestRepo)
	TestHandler := handlers.InitTestHandler(TestService)

	TestRouter.POST("/create", TestHandler.CreateTest)
	TestRouter.GET("/:id", TestHandler.GetTest)
	TestRouter.PUT("/add", TestHandler.AddTest)
	TestRouter.POST("/answer", TestHandler.TestAnswer)

	return TestRouter
}
