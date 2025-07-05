package main

import (
	"github.com/Claudio712005/go-task-api/config"
	"github.com/Claudio712005/go-task-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.CarregarConfiguracoes()

	router := gin.Default()

	api := router.Group("/api")

	routes.CarregarRotas(api)

	router.Run()
}