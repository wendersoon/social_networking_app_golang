package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configurar utiliza as variáveis de ambiente para a criação do SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar registra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	//condificar os dados para salvar no cookie
	dadosCodificados, err := s.Encode("dados", dados)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("dados")
	if err != nil {
		return nil, err
	}

	valores := make(map[string]string)

	if err := s.Decode("dados", cookie.Value, &valores); err != nil {
		return nil, err
	}

	return valores, nil
}

// Deletar limpa os dados codificados do cookie
func Deletar(w http.ResponseWriter) {

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})

}
