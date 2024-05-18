package delivery

import (
	"Tatarinhack_backend/docs"
	"Tatarinhack_backend/internal/delivery/handlers"
	"Tatarinhack_backend/internal/repository/user"
	userserv "Tatarinhack_backend/internal/service/user"
	"Tatarinhack_backend/internal/delivery/middleware"
	"Tatarinhack_backend/internal/delivery/routers"
	"Tatarinhack_backend/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, logger *logger.Logs) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	middlewareStruct := middleware.InitMiddleware(logger)
	r.Use(middlewareStruct.CORSMiddleware())

	userRouter := r.Group("/user")

	userRepo := user.InitUserRepo(db)
	userService := userserv.InitUserRepo(userRepo)
	userHandler := handlers.InitUserHandler(userService)

	userRouter.POST("/create", userHandler.Create)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/get/:id", userHandler.GetUser)
	userRouter.GET("/get/friend/:id", userHandler.GetFriend)
	userRouter.GET("/get/man/:id", userHandler.GetMan)
	userRouter.PUT("/pwd", userHandler.UpdatePassword)
	userRouter.PUT("/gram/:id", userHandler.GrammarUp)
	userRouter.PUT("/voc/:id", userHandler.VocabularyUp)
	userRouter.PUT("/speak/:id", userHandler.SpeakingUp)
	userRouter.POST("/add", userHandler.AddFriend)
	userRouter.DELETE("/delete/:id", userHandler.Delete)
	userRouter.GET("/get/friend/lst/:id", userHandler.FriendsList) // Friendslist
	//userRouter.PUT("/:id", userHandler.) // UpdAchievment
	//userRouter.PUT("/:id", userHandler.) // UpdDays
	userRouter.PUT("/lvl/:id", userHandler.LevelUp) // UpdLevel
\
	routers.InitRouting(r, db, logger, middlewareStruct)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
