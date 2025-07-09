package models

import (
	"errors"
	"strings"
	"time"

	"github.com/Claudio712005/go-task-api/util"
)

// Usuario representa um usuário do sistema
type Usuario struct {
	ID       uint64    `gorm:"primaryKey" json:"id"`
	Nome     string    `gorm:"type:varchar(100);not null" json:"nome" validate:"required,min=3"`     
	Email    string    `gorm:"type:varchar(100);not null;unique" json:"email" validate:"required,email"`
	Senha    string    `gorm:"type:varchar(200);not null" json:"senha" validate:"required,min=6"`
	CriadoEm time.Time `gorm:"autoCreateTime" json:"criado_em"`
}

// Validar valida os campos do usuário
func (u *Usuario) Validar(tipo string) error {
	if err := util.ValidarCampos(u); err != nil && tipo == "cadastrar" {
		return err
	}

	if(tipo == "atualizar") {
		if u.ID == 0 {
			return errors.New("ID do usuário não pode ser zero")
		}
		if u.Nome == "" {
			return errors.New("nome é obrigatório")
		}
		if u.Email == "" {
			return errors.New("email é obrigatório")
		}
	}

	u.preparar()
	return nil
}

// Preparar prepara o usuário antes de salvar no banco de dados
func (u *Usuario) preparar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
}
