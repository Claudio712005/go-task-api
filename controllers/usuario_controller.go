package controllers

import (
	"github.com/Claudio712005/go-task-api/config"
	"github.com/Claudio712005/go-task-api/models"
	"github.com/Claudio712005/go-task-api/repository"
	"github.com/Claudio712005/go-task-api/security"
	"github.com/Claudio712005/go-task-api/util"
	"github.com/gin-gonic/gin"
)

func CadastrarUsuarioHandler(c *gin.Context) {

	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		util.ResponseError(c, 400, "Dados inválidos")
		return
	}

	if err := usuario.Validar(); err != nil {
		util.ResponseError(c, 400, "Erro de validação: "+err.Error())
		return
	}
	
	senha, err := security.CriptografarSenha(usuario.Senha)
	if err != nil {
		util.ResponseError(c, 500, "Erro ao criptografar senha")
		return
	}

	usuario.Senha = senha

	repositorio := repository.NewUsuarioRepository(config.DB)

	usuarioExistente, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil && err.Error() != "record not found" {
		util.ResponseError(c, 500, "Erro ao verificar usuário existente")
		return
	}
	if usuarioExistente != nil {
		util.ResponseError(c, 409, "Usuário já cadastrado com este e-mail")
		return
	}

	id, err := repositorio.CadastrarUsuario(&usuario)
	if err != nil {
		util.ResponseError(c, 500, "Erro ao cadastrar usuário")
		return
	}

	util.ResponseSuccess(c, 200, id)
}
