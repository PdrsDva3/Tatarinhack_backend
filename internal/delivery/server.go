package delivery

import (
	"Tatarinhack_backend/docs"
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
	//r.SetTrustedProxies([]string{"127.0.0.1"})
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	middlewareStruct := middleware.InitMiddleware(logger)
	r.Use(middlewareStruct.CORSMiddleware())

	routers.InitRouting(r, db, logger, middlewareStruct)
	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
