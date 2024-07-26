package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto retorna um modelo Usuario com todas as informações preenchidas
// Essa função faz 4 requisições na API
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {

		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar o usuário")
			}
			usuario = usuarioCarregado
		case seguidoresCarregado := <-canalSeguidores:
			if seguidoresCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar seguidores")
			}
			seguidores = seguidoresCarregado

		case seguindoCarregado := <-canalSeguindo:
			if seguindoCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar que o usuário está seguindo")
			}
			seguindo = seguindoCarregado
		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao buscar as publicações")
			}
			publicacoes = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

// BuscarDadosDoUsuario chama a API para buscar os dados básicos do usuário
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioId)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- Usuario{}
	}
	defer response.Body.Close()

	var usuario Usuario
	if err := json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		canal <- Usuario{}
	}

	canal <- usuario

}

// BuscarSeguidores chama a API para buscar seguidores do usuário
func BuscarSeguidores(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioId)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if err := json.NewDecoder(response.Body).Decode(&seguidores); err != nil {
		canal <- nil
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

// BuscarSeguindo chama API para buscar os usuários seguidos por um usuário
func BuscarSeguindo(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioId)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if err := json.NewDecoder(response.Body).Decode(&seguindo); err != nil {
		canal <- nil
	}

	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

func BuscarPublicacoes(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioId)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if err := json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		canal <- nil
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
