package rotas

import (
	"net/http"

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

	for _, rota := range rotas {
		router.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}

	// caminho dos arquivos estáticos
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", fileServer))
	return router
}
