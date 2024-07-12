package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (u *Usuario) Preparar(etapa string) error {
	if err := u.validar(etapa); err != nil {
		return err
	}

	if err := u.formatar(etapa); err != nil {
		return err
	}

	return nil
}

func (u *Usuario) validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("o campo nome e obrigatorio")
	}

	if u.Nick == "" {
		return errors.New("o campo nick e obrigatorio")
	}

	if u.Email == "" {
		return errors.New("o campo email e obrigatorio")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("o email inserido e invalido")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("o campo senha e obrigatorio")
	}

	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(u.Senha)
		if err != nil {
			return err
		}

		u.Senha = string(senhaComHash)
	}
	return nil
}
