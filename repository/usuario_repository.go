package repository

import (
	"github.com/Claudio712005/go-task-api/models"
	"gorm.io/gorm"
)

// UsuarioRepository define os métodos que o repositório de usuários deve implementar
type UsuarioRepository interface {
	CadastrarUsuario(usuario *models.Usuario) (uint64, error)
	BuscarPorEmail(email string) (*models.Usuario, error)
}

type usuarioRepository struct {
	db *gorm.DB
}

// NewUsuarioRepository cria uma nova instância de UsuarioRepository
func NewUsuarioRepository(db *gorm.DB) UsuarioRepository {
	return &usuarioRepository{
		db: db,
	}
}

// CadastrarUsuario cadastra um novo usuário no banco de dados
func (r *usuarioRepository) CadastrarUsuario(usuario *models.Usuario) (uint64, error) {
	if err := r.db.Create(usuario).Error; err != nil {
		return 0, err
	}
	return usuario.ID, nil
}

// BuscarPorEmail busca um usuário pelo email
func (r *usuarioRepository) BuscarPorEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := r.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}