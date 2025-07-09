package routes

import (
	"github.com/Claudio712005/go-task-api/middleware"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
	hasAuth bool
}

// CarregarRotas carrega todas as rotas na instância do router
func CarregarRotas(router *gin.RouterGroup){
	rotas := []Route{}
	
	rotas = append(rotas, AgruparRotasUsuarios()...)
	rotas = append(rotas, AgruparRotasAutenticacao()...)
	rotas = append(rotas, AgruparRotasTarefas()...)

	for _, rota := range rotas {
		if rota.hasAuth {
			router.Handle(rota.Method, rota.Path, middleware.AutenticacaoMiddleware(), rota.Handler)
		} else {
			router.Handle(rota.Method, rota.Path, rota.Handler)
		}
	}
}