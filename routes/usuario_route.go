package routes

import (
	"github.com/Claudio712005/go-task-api/controllers"
)

func AgruparRotasUsuarios() []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    "/usuarios",
			Handler: controllers.CadastrarUsuarioHandler,
			hasAuth: false,
		},
		{
			Method:  "GET",
			Path:    "/usuarios/:id",
			Handler: controllers.BuscarUsuarioPorIdHandler,
			hasAuth: true,
		},
	}
}
