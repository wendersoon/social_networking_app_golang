package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarPublicação adiciona uma nova publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	usuarioId, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacao modelos.Publicacao
	if err := json.Unmarshal(corpoRequest, &publicacao); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	publicacao.AutorID = usuarioId

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)

	publicacao.ID, err = repositorio.Criar(publicacao)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
	}

	respostas.JSON(w, http.StatusCreated, publicacao)

}

// BuscarPublicacoes retorna as publicações que aparecem no feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

// BuscarPublicacao retorna apenas uma publicacao
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// AtualizarPublicacao edita as informações de uma publicação
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

// DeletarPublicacao remove uma publicação do banco de dados
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}