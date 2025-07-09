package repository

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Claudio712005/go-task-api/models"
	"gorm.io/gorm"
)

// TarefaRepository define os métodos que o repositório de tarefas deve implementar
type TarefaRepository interface {
	CadastrarTarefa(tarefa *models.Tarefa) (uint64, error)
	BuscarTarefasPorUsuario(usuarioID uint64) ([]models.Tarefa, error)
	BuscarTarefaPorTitulo(titulo string) (*models.Tarefa, error)
	BuscarTarefaPorID(id uint64) (*models.Tarefa, error)
	AtualizarTarefa(tarefa *models.Tarefa) error
	ConcluirTarefa(id uint64) error
	DeletarTarefa(id uint64) error
	BuscarTarefasPaginado(idUsuario uint64, pagina *models.Page) (error)
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

// BuscarTarefaPorID busca uma tarefa pelo ID
func (r *tarefaRepository) BuscarTarefaPorID(id uint64) (*models.Tarefa, error) {

	if id == 0 {
		return nil, errors.New("ID inválido")
	}

	var tarefa models.Tarefa
	if err := r.db.First(&tarefa, id).Error; err != nil {
		return nil, err
	}

	return &tarefa, nil
}

// AtualizarTarefa atualiza uma tarefa existente
func (r *tarefaRepository) AtualizarTarefa(tarefa *models.Tarefa) error {
	update := map[string]interface{}{
		"titulo":    tarefa.Titulo,
		"descricao": tarefa.Descricao,
	}

	if tarefa.ID == 0 {
		return gorm.ErrRecordNotFound
	}

	if err := r.db.Model(&tarefa).Updates(update).Error; err != nil {
		return err
	}

	return nil
}

// ConcluirTarefa marca uma tarefa como concluída
func (r *tarefaRepository) ConcluirTarefa(id uint64) error {
	if id == 0 {
		return errors.New("ID inválido")
	}

	update := map[string]interface{}{
		"concluida":    true,
		"concluida_em": time.Now(),
	}

	if err := r.db.Model(&models.Tarefa{}).Where("id = ?", id).Updates(update).Error; err != nil {
		return err
	}

	return nil
}

// DeletarTarefa deleta uma tarefa pelo ID
func (r *tarefaRepository) DeletarTarefa(id uint64) error {
	if id == 0 {
		return errors.New("ID inválido")
	}

	if err := r.db.Delete(&models.Tarefa{}, id).Error; err != nil {
		return err
	}

	return nil
}

// BuscarTarefasPaginado busca tarefas de um usuário com paginação
func (r *tarefaRepository) BuscarTarefasPaginado(idUsuario uint64, pagina *models.Page) (error) {
	var tarefas []models.Tarefa
	var total int64

	if idUsuario == 0 {
		return errors.New("ID de usuário inválido")
	}

	if pagina.Page == 0 {
		pagina.Page = 1
	}
	if pagina.Limit == 0 {
		pagina.Limit = 10
	}
	if pagina.SortBy == "" {
		pagina.SortBy = "criado_em"
	}
	if pagina.SortOrder == "" {
		pagina.SortOrder = "desc"
	}

	validSortFields := map[string]bool{
		"id":            true,
		"titulo":        true,
		"criado_em":     true,
		"atualizado_em": true,
		"concluida":     true,
	}

	sortField := strings.ToLower(pagina.SortBy)
	sortOrder := strings.ToLower(pagina.SortOrder)

	if !validSortFields[sortField] {
		sortField = "criado_em"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	orderClause := fmt.Sprintf("%s %s", sortField, sortOrder)

	if err := r.db.Model(&models.Tarefa{}).
		Where("usuario_id = ?", idUsuario).
		Count(&total).Error; err != nil {
		return err
	}

	if err := r.db.Where("usuario_id = ?", idUsuario).
		Order(orderClause).
		Limit(int(pagina.Limit)).
		Offset(int((pagina.Page - 1) * pagina.Limit)).
		Find(&tarefas).Error; err != nil {
		return err
	}

	pagina.Total = total
	pagina.TotalPages = (pagina.Total + pagina.Limit - 1) / pagina.Limit

	content := make([]interface{}, len(tarefas))

	for i, tarefa := range tarefas {
		content[i] = tarefa
	}

	pagina.Content = content
	return nil
}
