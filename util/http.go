package util

import "github.com/gin-gonic/gin"

// ResponseError envia uma resposta de erro para o cliente
func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"success": false,
		"message": message,
	})
}

// ResponseSuccess envia uma resposta de sucesso para o cliente
func ResponseSuccess(c *gin.Context, status int, data interface{}){
	c.JSON(status, gin.H{
		"success": true,
		"data":    data,
	})
}