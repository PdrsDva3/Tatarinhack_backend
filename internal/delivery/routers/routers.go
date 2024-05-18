package routers

import (
	"Tatarinhack_backend/internal/delivery/middleware"
	"Tatarinhack_backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouting(r *gin.Engine, db *sqlx.DB, logger *logger.Logs, middlewareStruct middleware.Middleware) {
	_ = RegisterTeachRouter(r, db, logger)
	_ = RegisterAnswerRouter(r, db, logger)
	_ = RegisterQuestionRouter(r, db, logger)
	_ = RegisterTestRouter(r, db, logger)
}
