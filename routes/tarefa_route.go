package routes

import "github.com/Claudio712005/go-task-api/controllers"

const TAREFA_BASE_PATH = "/tarefas"

func tarefaBasePath(path string) string {
	return TAREFA_BASE_PATH + path
}

// AgruparRotasTarefas agrupa as rotas relacionadas a tarefas
func AgruparRotasTarefas() []Route {
	return []Route{
		{
			Method:  "POST",
			hasAuth: true,
			Path:    tarefaBasePath(""),
			Handler: controllers.CadastrarTarefaHandler,
		},
		{
			Method:  "GET",
			hasAuth: true,
			Path:   tarefaBasePath(""),
			Handler: controllers.BuscarTarefasPorUsuarioHandler,
		},
	}
}
