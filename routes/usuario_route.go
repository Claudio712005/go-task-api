package routes

import (
	"fmt"

	"github.com/Claudio712005/go-task-api/controllers"
)

const USUARIO_BASE_PATH = "/usuarios"

func usuarioBasePath(path string) string {
	return fmt.Sprintf("%s%s", USUARIO_BASE_PATH, path)
}

// AgruparRotasUsuarios agrupa as rotas relacionadas a usuários
func AgruparRotasUsuarios() []Route {
	return []Route{
		{
			Method:  "POST",
			Path:    usuarioBasePath(""),
			Handler: controllers.CadastrarUsuarioHandler,
			hasAuth: false,
		},
		{
			Method:  "GET",
			Path:    usuarioBasePath("/:id"),
			Handler: controllers.BuscarUsuarioPorIdHandler,
			hasAuth: true,
		},
		{
			Method:  "PUT",
			Path:    usuarioBasePath("/:id"),
			Handler: controllers.AtualizarUsuarioHandler,
			hasAuth: true,
		},
		{
			Method:  "DELETE",
			Path:    usuarioBasePath("/:id"),
			Handler: controllers.DeletarUsuarioHandler,
			hasAuth: true,
		},
		{
			Method: "POST",
			Path:  usuarioBasePath("/senha"),
			Handler: controllers.AtualizarSenhaHandler,
			hasAuth: true,
		},
	}
}
