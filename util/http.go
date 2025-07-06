package util

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"message" example:"Erro"`
	Success bool   `json:"success" example:"false"`
}

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// ResponseError envia uma resposta de erro para o cliente
func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, ErrorResponse{
		message,
		false,
	})
}

// ResponseSuccess envia uma resposta de sucesso para o cliente
func ResponseSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, SuccessResponse{
		data,
		true,
	})
}
