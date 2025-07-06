package controllers

import (
	"github.com/Claudio712005/go-task-api/config"
	"github.com/Claudio712005/go-task-api/models"
	"github.com/Claudio712005/go-task-api/repository"
	"github.com/Claudio712005/go-task-api/security"
	"github.com/Claudio712005/go-task-api/util"
	"github.com/gin-gonic/gin"
)

// AutenticarUsuarioHandler godoc
// @Summary Autentica um usuário
// @Description Autentica um usuário no sistema e retorna um token JWT
// @Tags Autenticação
// @Accept json
// @Produce json
// @Param login body models.Login true "Dados de login do usuário"
// @Success 200 {object} util.SuccessResponse "Usuário autenticado com sucesso"
// @Failure 400 {object} util.ErrorResponse "Dados inválidos"
// @Failure 401 {object} util.ErrorResponse "Senha incorreta"
// @Failure 404 {object} util.ErrorResponse "Usuário não encontrado"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /autenticar [post]
// AutenticarUsuarioHandler autentica um usuário e retorna um token JWT
func AutenticarUsuarioHandler(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		util.ResponseError(c, 400, "Dados inválidos")
		return
	}

	if err := login.Validar(); err != nil {
		util.ResponseError(c, 400, "Erro de validação: "+err.Error())
		return
	}

	repositorio := repository.NewUsuarioRepository(config.DB)
	usuario, err := repositorio.BuscarPorEmail(login.Email)
	if err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Usuário não encontrado")
			return
		}
		util.ResponseError(c, 500, "Erro ao buscar usuário")
		return
	}

	if err := security.VerificarSenha(login.Senha, usuario.Senha); err != nil {
		util.ResponseError(c, 401, "Senha incorreta")
		return
	}

	token, err := security.GerarToken(usuario.ID)
	if err != nil {
		util.ResponseError(c, 500, "Erro ao gerar token")
		return
	}

	usuario.Senha = ""

	util.ResponseSuccess(c, 200, gin.H{
		"token":   token,
		"usuario": usuario,
	})
}
