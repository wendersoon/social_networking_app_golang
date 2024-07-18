package modelos

// DadosAuth contem o token e o id do usuário autenticado
type DadosAuth struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
