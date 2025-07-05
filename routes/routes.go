package routes

import "github.com/gin-gonic/gin"

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

	for _, rota := range rotas {
		if rota.hasAuth {
			// ...
		} else {
			router.Handle(rota.Method, rota.Path, rota.Handler)
		}
	}
}