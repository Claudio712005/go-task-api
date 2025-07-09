package models

import "github.com/Claudio712005/go-task-api/util"

type Senha struct {
	SenhaNova  string `json:"senha_nova" validate:"required,min=6"`
	SenhaAtual string `json:"senha_atual" validate:"required,min=6"`
}

func (s *Senha) Validar(acao string) error {

	if err := util.ValidarCampos(s); err != nil {
		return err
	}

	return nil
}