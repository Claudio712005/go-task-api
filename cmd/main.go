package main

import (
	"github.com/Claudio712005/go-task-api/config"
	"github.com/Claudio712005/go-task-api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/Claudio712005/go-task-api/docs"
)

// @title Go Task API
// @version 1.0
// @description Esta é a API de gerenciamento de tarefas
// @host localhost:8080
// @BasePath /api
func main() {

	config.CarregarConfiguracoes()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	routes.CarregarRotas(api)

	router.Run()
}