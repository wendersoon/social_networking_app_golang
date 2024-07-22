package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasLogout = Rota{
	Uri:                "/logout",
	Metodo:             http.MethodGet,
	Funcao:             controllers.FazerLogout,
	RequerAutenticacao: true,
}
