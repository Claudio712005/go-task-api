package controllers

import (
	"strconv"

	"github.com/Claudio712005/go-task-api/config"
	"github.com/Claudio712005/go-task-api/models"
	"github.com/Claudio712005/go-task-api/repository"
	"github.com/Claudio712005/go-task-api/security"
	"github.com/Claudio712005/go-task-api/util"
	"github.com/gin-gonic/gin"
)

// CadastrarUsuarioHandler godoc
// @Summary Cadastra um novo usuário
// @Description Cadastra um novo usuário no sistema
// @Tags Usuários
// @Accept json
// @Produce json
// @Param usuario body models.Usuario true "Dados do usuário"
// @Success 201 {object} util.SuccessResponse "Usuário cadastrado com sucesso"
// @Failure 400 {object} util.ErrorResponse "Dados inválidos"
// @Failure 409 {object} util.ErrorResponse "Usuário já cadastrado"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /usuarios [post]
// CadastrarUsuarioHandler cadastra um novo usuário
func CadastrarUsuarioHandler(c *gin.Context) {

	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		util.ResponseError(c, 400, "Dados inválidos")
		return
	}

	if err := usuario.Validar("cadastrar"); err != nil {
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

	util.ResponseSuccess(c, 201, id)
}

// BuscarUsuarioPorIdHandler godoc
// @Summary Busca um usuário pelo ID
// @Description Busca um usuário pelo ID fornecido
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path uint64 true "ID do usuário"
// @Success 200 {object} util.SuccessResponse "Usuário encontrado"
// @Failure 404 {object} util.ErrorResponse "Usuário não encontrado"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /usuarios/{id} [get]
// BuscarUsuarioPorIdHandler busca um usuário pelo ID
func BuscarUsuarioPorIdHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		util.ResponseError(c, 400, "ID inválido")
		return
	}

	repositorio := repository.NewUsuarioRepository(config.DB)
	usuario, err := repositorio.BuscarPorID(id)
	if err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Usuário não encontrado")
			return
		}
		util.ResponseError(c, 500, "Erro ao buscar usuário")
		return
	}

	util.ResponseSuccess(c, 200, usuario)
}

// AtualizarUsuarioHandler godoc
// @Summary Atualiza um usuário existente
// @Description Atualiza os dados de um usuário existente
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path uint64 true "ID do usuário"
// @Param usuario body models.Usuario true "Dados do usuário"
// @Success 200 {object} util.SuccessResponse "Usuário atualizado com sucesso"
// @Failure 400 {object} util.ErrorResponse "Dados inválidos"
// @Failure 403 {object} util.ErrorResponse "Você não tem permissão para atualizar
// @Failure 404 {object} util.ErrorResponse "Usuário não encontrado"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /usuarios/{id} [put]
// AtualizarUsuarioHandler atualiza os dados de um usuário existente
func AtualizarUsuarioHandler(c *gin.Context) {

	var request *models.Usuario

	if err := c.ShouldBindJSON(&request); err != nil {
		util.ResponseError(c, 400, "Corpo da requisição inválido")
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		util.ResponseError(c, 400, "ID inválido")
		return
	}

	tokenId, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Token inválido")
		return
	}

	if id != tokenId {
		util.ResponseError(c, 403, "Você não tem permissão para atualizar este usuário")
		return
	}

	request.ID = id
	if err := request.Validar("atualizar"); err != nil {
		util.ResponseError(c, 400, "Erro de validação: "+err.Error())
		return
	}

	repositorio := repository.NewUsuarioRepository(config.DB)

	usuarioExistePorEmail, err := repositorio.BuscarPorEmail(request.Email)
	if err != nil {
		util.ResponseError(c, 500, "Erro ao verificar usuário existente")
		return
	}

	if usuarioExistePorEmail != nil && usuarioExistePorEmail.ID != request.ID {
		util.ResponseError(c, 409, "Já existe um usuário cadastrado com este e-mail")
		return
	}

	if err := repositorio.AtualizarUsuario(request); err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Usuário não encontrado")
			return
		}
		util.ResponseError(c, 500, "Erro ao atualizar usuário")
		return
	}

	util.ResponseSuccess(c, 200, "Usuário atualizado com sucesso")
}

// DeletarUsuarioHandler godoc
// @Summary Deleta um usuário existente
// @Description Deleta um usuário existente pelo ID
// @Tags Usuários
// @Accept json
// @Produce json
// @Param id path uint64 true "ID do usuário"
// @Success 204 "Usuário deletado com sucesso"
// @Failure 400 {object} util.ErrorResponse "ID inválido"
// @Failure 401 {object} util.ErrorResponse "Token inválido"
// @Failure 403 {object} util.ErrorResponse "Você não tem permissão para deletar este usuário"
// @Failure 404 {object} util.ErrorResponse "Usuário não encontrado"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /usuarios/{id} [delete]
// DeletarUsuarioHandler deleta um usuário existente pelo ID
func DeletarUsuarioHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		util.ResponseError(c, 400, "ID inválido")
		return
	}

	idToken, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Token inválido")
		return
	}

	if id != idToken {
		util.ResponseError(c, 403, "Você não tem permissão para deletar este usuário")
		return
	}

	repositorio := repository.NewUsuarioRepository(config.DB)
	if err := repositorio.DeletarUsuario(id); err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Usuário não encontrado")
			return
		}
		util.ResponseError(c, 500, "Erro ao deletar usuário")
		return
	}

	util.ResponseSuccess(c, 204, nil)
}