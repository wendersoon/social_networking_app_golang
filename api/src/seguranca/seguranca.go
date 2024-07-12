package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e colaca uma hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha com hash e retorna se elas são iguais
func VerificarSenha(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
