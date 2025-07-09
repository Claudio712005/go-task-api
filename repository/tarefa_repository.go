package repository

import (
	"errors"

	"github.com/Claudio712005/go-task-api/models"
	"gorm.io/gorm"
)

// TarefaRepository define os métodos que o repositório de tarefas deve implementar
type TarefaRepository interface {
	CadastrarTarefa(tarefa *models.Tarefa) (uint64, error)
	BuscarTarefasPorUsuario(usuarioID uint64) ([]models.Tarefa, error)
	BuscarTarefaPorTitulo(titulo string) (*models.Tarefa, error)
}

type tarefaRepository struct {
	db *gorm.DB
}

// NewTarefaRepository cria uma nova instância de TarefaRepository
func NewTarefaRepository(db *gorm.DB) TarefaRepository {
	return &tarefaRepository{
		db: db,
	}
}

// CadastrarTarefa cadastra uma nova tarefa no banco de dados
func (r *tarefaRepository) CadastrarTarefa(tarefa *models.Tarefa) (uint64, error) {
	if tarefa.UsuarioID == 0 {
		return 0, errors.New("usuário não informado")
	}

	if err := r.db.Create(tarefa).Error; err != nil {
		return 0, err
	}

	return tarefa.ID, nil
}

// BuscarTarefasPorUsuario busca todas as tarefas de um usuário
func (r *tarefaRepository) BuscarTarefasPorUsuario(usuarioID uint64) ([]models.Tarefa, error) {
	var tarefas []models.Tarefa
	if err := r.db.Where("usuario_id = ?", usuarioID).Find(&tarefas).Error; err != nil {
		return nil, err
	}
	
	return tarefas, nil
}

// BuscarTarefaPorTitulo busca uma tarefa pelo título
func (r *tarefaRepository) BuscarTarefaPorTitulo(titulo string) (*models.Tarefa, error) {
	var tarefa models.Tarefa
	if err := r.db.Where("LOWER(titulo) = LOWER(?)", titulo).First(&tarefa).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("tarefa não encontrada")
		}
		return nil, err
	}
	return &tarefa, nil
}