package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da aplicação
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar organiza todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaHome)
	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			router.HandleFunc(rota.Uri,
				middlewares.Logger(
					middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)

		} else {
			router.HandleFunc(rota.Uri,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}

	}

	// caminho dos arquivos estáticos
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", fileServer))
	return router
}
