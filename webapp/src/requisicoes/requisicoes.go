package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutenticacao é usado para setar o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(metodo, url, dados)
	if err != nil {
		return nil, err
	}

	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
