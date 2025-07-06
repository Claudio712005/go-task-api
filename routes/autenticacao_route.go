package routes

import "github.com/Claudio712005/go-task-api/controllers"

func AgruparRotasAutenticacao() []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    "/autenticar",
			Handler: controllers.AutenticarUsuarioHandler,
			hasAuth: false,
		},
	}
}
