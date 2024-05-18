package routers

import (
	"Tatarinhack_backend/internal/delivery/handlers"
	"Tatarinhack_backend/internal/repository/course"
	courseserv "Tatarinhack_backend/internal/service/course"
	"Tatarinhack_backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterCourseRouter(r *gin.Engine, db *sqlx.DB, logs *logger.Logs) *gin.RouterGroup {

	CourseRouter := r.Group("/course")

	CourseRepo := course.InitCourseRepository(db)
	CourseService := courseserv.InitCourseService(CourseRepo)
	CourseHandler := handlers.InitCourseHandler(CourseService)

	CourseRouter.POST("/create", CourseHandler.CreateCourse)
	CourseRouter.GET("/:id", CourseHandler.GetCourse)
	CourseRouter.PUT("/add", CourseHandler.AddCourse)
	return CourseRouter
}
