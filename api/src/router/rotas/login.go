package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Uri:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
