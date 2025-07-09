package models

import (
	"errors"
	"strings"
	"time"

	"github.com/Claudio712005/go-task-api/util"
)

// Tarefa representa uma tarefa no sistema
type Tarefa struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	Titulo       string    `json:"titulo" gorm:"type:varchar(100);not null" validate:"required,min=3"`
	Descricao    string    `json:"descricao" gorm:"type:text" validate:"required,min=5"`
	UsuarioID    uint64    `json:"usuario_id" gorm:"not null" validate:"required"`
	CriadoEm     time.Time `json:"criado_em" gorm:"autoCreateTime"`
	AtualizadoEm time.Time `json:"atualizado_em" gorm:"autoUpdateTime"`
	Concluida    bool      `json:"concluida" gorm:"default:false"`
	ConcluidaEm  *time.Time `json:"concluida_em,omitempty"`
}

// Validar valida os campos da tarefa
func (t *Tarefa) Validar(tipo string) error {
	
	if err := util.ValidarCampos(t); err != nil && tipo == "cadastrar" {
		return err
	}

	if tipo == "atualizar" {
		if t.ID == 0 {
			return errors.New("ID da tarefa não pode ser zero")
		}
		if t.Titulo == "" {
			return errors.New("título é obrigatório")
		}
		if t.Descricao == "" {
			return errors.New("descrição é obrigatória")
		}
	}

	t.preparar()

	return nil
}

func (t *Tarefa) preparar(){
	t.Titulo = strings.TrimSpace(t.Titulo)
	t.Descricao = strings.TrimSpace(t.Descricao)
}