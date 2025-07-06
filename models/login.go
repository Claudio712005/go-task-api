package models

import "github.com/Claudio712005/go-task-api/util"

// Login representa os dados de login do usuário
type Login struct {
	Email	string `json:"email" validate:"required,email"`
	Senha	string `json:"senha" validate:"required,min=6"`
}

func (l *Login) Validar() error {
	if err := util.ValidarCampos(l); err != nil {
		return err
	}

	return nil
}