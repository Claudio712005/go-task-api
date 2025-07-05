package models

import (
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
	CriadoEm time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// Validar valida os campos do usuário
func (u *Usuario) Validar() error {
	if err := util.ValidarCampos(u); err != nil {
		return err
	}

	u.preparar()
	return nil
}

// Preparar prepara o usuário antes de salvar no banco de dados
func (u *Usuario) preparar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Email = strings.TrimSpace(u.Email)
}
