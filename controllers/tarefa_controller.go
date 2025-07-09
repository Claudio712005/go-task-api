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

	if err := tarefa.Validar("cadastrar"); err != nil {
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

// BuscarTarefaPorIDHandler godoc
// @Summary Busca uma tarefa por ID
// @Description Busca uma tarefa específica pelo ID fornecido
// @Tags Tarefas
// @Accept json
// @Produce json
// @Param id path uint64 true "ID da tarefa"
// @Success 200 {object} models.Tarefa "Tarefa encontrada"
// @Failure 400 {object} util.ErrorResponse "ID inválido"
// @Failure 401 {object} util.ErrorResponse "Usuário não autenticado"
// @Failure 403 {object} util.ErrorResponse "Usuário não autorizado"
// @Failure 404 {object} util.ErrorResponse "Tarefa não encontrada"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /tarefas/{id} [get]
// BuscarTarefaPorIDHandler busca uma tarefa pelo ID fornecido
func BuscarTarefaPorIDHandler(c *gin.Context) {
	idTarefa := c.Param("id")
	id, err := strconv.ParseUint(idTarefa, 10, 64)
	if err != nil {
		util.ResponseError(c, 400, "ID inválido")
		return
	}

	repositorio := repository.NewTarefaRepository(config.DB)
	tarefa, err := repositorio.BuscarTarefaPorID(id)
	if err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Tarefa não encontrada")
			return
		}
		util.ResponseError(c, 500, "Erro ao buscar tarefa: "+err.Error())
		return
	}

	idToken, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Usuário não autenticado")
		return
	}

	if tarefa.UsuarioID != idToken {
		util.ResponseError(c, 403, "Usuário não autorizado a visualizar esta tarefa")
		return
	}

	util.ResponseSuccess(c, 200, tarefa)
}

// AtualizarTarefaHandler godoc
// @Summary Atualiza uma tarefa existente
// @Description Atualiza os dados de uma tarefa existente
// @Tags Tarefas
// @Accept json
// @Produce json
// @Param id path uint64 true "ID da tarefa"
// @Param tarefa body models.Tarefa true "Dados da tarefa"
// @Success 200 {object} util.SuccessResponse "Tarefa atualizada com sucesso"
// @Failure 400 {object} util.ErrorResponse "Dados inválidos"
// @Failure 401 {object} util.ErrorResponse "Usuário não autenticado"
// @Failure 403 {object} util.ErrorResponse "Usuário não autorizado"
// @Failure 404 {object} util.ErrorResponse "Tarefa não encontrada"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /tarefas/{id} [put]
// AtualizarTarefaHandler atualiza os dados de uma tarefa existente
func AtualizarTarefaHandler(c *gin.Context) {
	idTarefa := c.Param("id")
	id, err := strconv.ParseUint(idTarefa, 10, 64)
	if err != nil {
		util.ResponseError(c, 400, "ID inválido")
		return
	}

	repositorio := repository.NewTarefaRepository(config.DB)
	tarefa, err := repositorio.BuscarTarefaPorID(id)
	if err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Tarefa não encontrada")
			return
		}
		util.ResponseError(c, 500, "Erro ao buscar tarefa: "+err.Error())
		return
	}

	idToken, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Usuário não autenticado")
		return
	}

	if tarefa.UsuarioID != idToken {
		util.ResponseError(c, 403, "Usuário não autorizado a atualizar esta tarefa")
		return
	}

	if err := c.ShouldBindJSON(&tarefa); err != nil {
		util.ResponseError(c, 400, "Dados inválidos")
		return
	}

	tarefa.ID = id

	if err := tarefa.Validar("atualizar"); err != nil {
		util.ResponseError(c, 400, "Erro de validação: "+err.Error())
		return
	}

	if err := repositorio.AtualizarTarefa(tarefa); err != nil {
		util.ResponseError(c, 500, "Erro ao atualizar tarefa: "+err.Error())
		return
	}

	util.ResponseSuccess(c, 200, "Tarefa atualizada com sucesso")
}

// ConcluirTarefaHandler godoc
// @Summary Conclui uma tarefa
// @Description Marca uma tarefa como concluída
// @Tags Tarefas
// @Accept json
// @Produce json
// @Param id path uint64 true "ID da tarefa"
// @Success 204 {object} util.SuccessResponse "Tarefa concluída com sucesso"
// @Failure 400 {object} util.ErrorResponse "ID inválido ou tarefa já concluída"
// @Failure 401 {object} util.ErrorResponse "Usuário não autenticado"
// @Failure 403 {object} util.ErrorResponse "Usuário não autorizado a concluir a tarefa de outro usuário"
// @Failure 404 {object} util.ErrorResponse "Tarefa não encontrada"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /tarefas/{id}/concluir [post]
// ConcluirTarefaHandler marca uma tarefa como concluída
func ConcluirTarefaHandler(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		util.ResponseError(c, 400, "ID inválido")
		return
	}

	repositorio := repository.NewTarefaRepository(config.DB)

	tarefa, err := repositorio.BuscarTarefaPorID(idUint)
	if err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Tarefa não encontrada")
			return
		}
		util.ResponseError(c, 500, "Erro ao buscar tarefa: "+err.Error())
		return
	}

	tokenID, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Usuário não autenticado")
		return
	}

	if tarefa.UsuarioID != tokenID {
		util.ResponseError(c, 403, "Usuário não autorizado a concluir a tarefa de outro usuário")
		return
	}

	if tarefa.Concluida {
		util.ResponseError(c, 400, "Tarefa já está concluída")
		return
	}

	if err := repositorio.ConcluirTarefa(idUint); err != nil {
		util.ResponseError(c, 500, "Erro ao concluir tarefa: "+err.Error())
		return
	}

	util.ResponseSuccess(c, 204, "")
}

// DeletarTarefa godoc
// @Summary Deleta uma tarefa
// @Description Deleta uma tarefa pelo ID
// @Tags Tarefas
// @Accept json
// @Produce json
// @Param id path uint64 true "ID da tarefa"
// @Success 204 {object} util.SuccessResponse "Tarefa deletada com sucesso"
// @Failure 400 {object} util.ErrorResponse "ID inválido"
// @Failure 401 {object} util.ErrorResponse "Usuário não autenticado"
// @Failure 403 {object} util.ErrorResponse "Usuário não autorizado a deletar esta tarefa"
// @Failure 404 {object} util.ErrorResponse "Tarefa não encontrada"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /tarefas/{id} [delete]
// DeletarTarefa deleta uma tarefa pelo ID
func DeletarTarefaHandler(c *gin.Context) {
	idTarefa := c.Param("id")
	id, err := strconv.ParseUint(idTarefa, 10, 64)
	if err != nil {
		util.ResponseError(c, 400, "ID inválido")
		return
	}

	tokenID, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Usuário não autenticado")
		return
	}

	repositorio := repository.NewTarefaRepository(config.DB)

	tarefaExistente, err := repositorio.BuscarTarefaPorID(id)
	if err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Tarefa não encontrada")
			return
		}
		util.ResponseError(c, 500, "Erro ao buscar tarefa: "+err.Error())
		return
	}

	if tarefaExistente.UsuarioID != tokenID {
		util.ResponseError(c, 403, "Usuário não autorizado a deletar esta tarefa")
		return
	}

	if err := repositorio.DeletarTarefa(id); err != nil {
		if err.Error() == "record not found" {
			util.ResponseError(c, 404, "Tarefa não encontrada")
			return
		}
		util.ResponseError(c, 500, "Erro ao deletar tarefa: "+err.Error())
		return
	}

	util.ResponseSuccess(c, 204, nil)
}

// BuscarTarefasPaginadoHandler godoc
// @Summary Busca tarefas paginadas
// @Description Busca tarefas de um usuário com paginação
// @Tags Tarefas
// @Accept json
// @Produce json
// @Param page query int false "Número da página" default(1)
// @Param limit query int false "Número de itens por página" default(10)
// @Param sort_by query string false "Campo para ordenação" default("id")
// @Param sort_order query string false "Direção da ordenação" default("asc")
// @Success 200 {object} models.Page "Página de tarefas"
// @Failure 400 {object} util.ErrorResponse "Parâmetros inválidos"
// @Failure 401 {object} util.ErrorResponse "Usuário não autenticado"
// @Failure 500 {object} util.ErrorResponse "Erro interno"
// @Router /tarefas/paginado [get]
// BuscarTarefasPaginadoHandler busca tarefas de um usuário com paginação
func BuscarTarefasPaginadoHandler(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	sortBy := c.DefaultQuery("sort_by", "id")
	sortOrder := c.DefaultQuery("sort_order", "asc")

	tokenID, err := security.ExtrairUsuarioID(c.GetHeader("Authorization"))
	if err != nil {
		util.ResponseError(c, 401, "Usuário não autenticado")
		return
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil || pageInt < 1 {
		util.ResponseError(c, 400, "Página inválida")
		return
	}

	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil || limitInt < 1 {
		util.ResponseError(c, 400, "Limite inválido")
		return
	}

	pagina := &models.Page{
		Page:      pageInt,
		Content:      []interface{}{},
		Limit:     limitInt,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	repositorio := repository.NewTarefaRepository(config.DB)

	if err := repositorio.BuscarTarefasPaginado(tokenID, pagina); err != nil {
		util.ResponseError(c, 500, "Erro ao buscar tarefas: "+err.Error())
		return
	}

	util.ResponseSuccess(c, 200, pagina)
}
