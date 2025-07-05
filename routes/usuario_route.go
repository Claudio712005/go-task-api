package routes

import (
	"github.com/Claudio712005/go-task-api/controllers"
)

func cadastrarUsuario() Route {
	return Route{
		Method:  "POST",
		Path:    "/usuarios",
		Handler: controllers.CadastrarUsuarioHandler,
		hasAuth: false,
	}
}

func AgruparRotasUsuarios() []Route {
	return []Route{
		cadastrarUsuario(),
	}
}
