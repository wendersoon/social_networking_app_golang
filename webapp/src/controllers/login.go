package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

// FazerLogin autentica o usuário na api
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutheticacao modelos.DadosAuth
	if err := json.NewDecoder(response.Body).Decode(&dadosAutheticacao); err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
	}

	w.Write([]byte("OK"))
}
