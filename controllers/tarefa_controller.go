package controllers

import (
	"github.com/Claudio712005/go-task-api/config"
	"github.com/Claudio712005/go-task-api/models"
	"github.com/Claudio712005/go-task-api/repository"
	"github.com/Claudio712005/go-task-api/security"
	"github.com/Claudio712005/go-task-api/util"
	"github.com/gin-gonic/gin"
)

// CadastrarTarefaHandler godoc
// @Summary Cadastra uma nova tarefa
// @Description Cadastra uma nova tarefa para o usuário autenticado
// @Tags Tarefas
// @Accept json
// @Produce json
// @Param tarefa body models.Tarefa true "Dados da tarefa"
// @Success 201 {object} util.SuccessResponse "Tarefa cadastrada com sucesso"
// @Failure 400 {object} util.ErrorResponse "Dados inválidos"
// @Failure 401 {object} util.ErrorResponse "Usuário não autenticado"
// @Failure 403 {object} util.ErrorResponse "Usuário não autorizado"
// @Failure 409 {object} util.ErrorResponse "Tarefa já cadastrada"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /tarefas [post]
// CadastrarTarefaHandler cadastra uma nova tarefa para o usuário autenticado
func CadastrarTarefaHandler(c *gin.Context) {
	var tarefa models.Tarefa

	if err := c.ShouldBindJSON(&tarefa); err != nil {
		util.ResponseError(c, 400, "Dados inválidos")
		return
	}

	if err := tarefa.Validar(); err != nil {
		util.ResponseError(c, 400, "Erro de validação: "+err.Error())
		return
	}

	idToken, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Usuário não autenticado")
		return
	}

	if idToken != tarefa.UsuarioID {
		util.ResponseError(c, 403, "Usuário não autorizado a criar uma tarefa para outro usuário")
		return
	}

	repositorio := repository.NewTarefaRepository(config.DB)

	if _, err := repositorio.BuscarTarefaPorTitulo(tarefa.Titulo); err == nil {
		util.ResponseError(c, 409, "Já existe uma tarefa com este título")
		return
	}

	tarefa.ID, err = repositorio.CadastrarTarefa(&tarefa)
	if err != nil {
		util.ResponseError(c, 500, "Erro ao cadastrar tarefa: "+err.Error())
		return
	}

	util.ResponseSuccess(c, 201, tarefa.ID)
}

// BuscarTarefasPorUsuarioHandler godoc
// @Summary Busca as tarefas de um usuário
// @Description Busca todas as tarefas associadas ao usuário autenticado
// @Tags Tarefas
// @Accept json
// @Produce json
// @Success 200 {array} models.Tarefa "Lista de tarefas do usuário"
// @Failure 401 {object} util.ErrorResponse "Usuário não autenticado"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /tarefas [get]
// BuscarTarefasPorUsuarioHandler busca as tarefas de um usuário autenticado
func BuscarTarefasPorUsuarioHandler(c *gin.Context) {
	tokenID, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Usuário não autenticado")
		return
	}

	repositorio := repository.NewTarefaRepository(config.DB)
	tarefas, err := repositorio.BuscarTarefasPorUsuario(tokenID)
	if err != nil {
		util.ResponseError(c, 500, "Erro ao buscar tarefas: "+err.Error())
		return
	}

	util.ResponseSuccess(c, 200, tarefas)
}
