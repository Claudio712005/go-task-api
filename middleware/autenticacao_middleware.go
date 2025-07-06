package middleware

import (
	"github.com/Claudio712005/go-task-api/security"
	"github.com/Claudio712005/go-task-api/util"
	"github.com/gin-gonic/gin"
)

// AutenticacaoMiddleware é um middleware que verifica se o token de autenticação está presente e é válido
func AutenticacaoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenBearer := c.GetHeader("Authorization")

		if tokenBearer == "" {
			c.AbortWithStatusJSON(401, util.ErrorResponse{
				Message: "Token de autenticação não fornecido",
				Success: false,
			})
			return
		}

		token := tokenBearer[len("Bearer "):]

		if err := security.ValidarToken(token); err != nil {
			c.AbortWithStatusJSON(401, util.ErrorResponse{
				Message: "Token de autenticação inválido: " + err.Error(),
				Success: false,
			})
			return
		}
	}
}
