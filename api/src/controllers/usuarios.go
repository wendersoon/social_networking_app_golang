package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario adiciona um usuário no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Modelos
	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := usuario.Preparar("cadastro"); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	// Banco
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Repositório
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, err = repositorio.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	//Respostas
	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	//Obtém a query referente ao usuario
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuarios, err := repositorio.Buscar(nomeOuNick)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusOK, usuarios)

}

// BuscarUsuario busca apenas um usuário salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	usuario, err := repositorio.BuscarPorID(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
	}

	respostas.JSON(w, http.StatusOK, usuario)

}

// AtualizarUsuario edita as informações de um usuário salvo no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadGateway, err)
		return
	}

	usuarioIdNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não é o seu."))
		return
	}

	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("edicao"); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if err = repositorio.Atualizar(usuarioID, usuario); err != nil {
		respostas.Erro(w, http.StatusBadGateway, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

// DeletarUsuario apaga um usuário salvo no banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadGateway, err)
		return
	}

	usuarioIdNoToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioID != usuarioIdNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("não é possível deletar usuário que não seja o seu"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	if err = repositorio.Deletar(usuarioID); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
