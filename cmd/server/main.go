package main

import (
	"community-qa-platform/internal/config"
	"community-qa-platform/internal/flags"
	"community-qa-platform/internal/handlers"
	"community-qa-platform/internal/repository"
	"community-qa-platform/internal/service"
	"community-qa-platform/pkg/database"
	"community-qa-platform/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	initProject()

	answerService := service.NewAnswerService(repository.NewAnswerRepo(database.DB))

	r := gin.Default()

	setupRoutes(r, answerService)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func initProject() {
	flag := flags.LoadFlags()
	conf := config.InitConfig(flag.Config)
	database.InitDB(conf.Db)
}

// setupRoutes 设置应用的路由
func setupRoutes(router *gin.Engine, as *service.AnswerService) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎来到社区问答平台")
	})

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	api := router.Group("/api", handlers.JwtMiddleware(), middleware.Logger())

	answerHandler := handlers.NewAnswerHandler(as)
	api.POST("/answers", answerHandler.CreateAnswer)
	api.GET("/answers/:id", answerHandler.GetAnswer)
}
