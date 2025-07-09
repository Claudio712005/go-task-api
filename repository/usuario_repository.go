package repository

import (
	"github.com/Claudio712005/go-task-api/models"
	"gorm.io/gorm"
)

// UsuarioRepository define os métodos que o repositório de usuários deve implementar
type UsuarioRepository interface {
	CadastrarUsuario(usuario *models.Usuario) (uint64, error)
	BuscarPorEmail(email string) (*models.Usuario, error)
	BuscarPorID(id uint64) (*models.Usuario, error)
	BuscarSenha(id uint64) (string, error)
	AtualizarUsuario(usuario *models.Usuario) error
	DeletarUsuario(id uint64) error
	AtualizarSenha(id uint64, senhaHash string) error
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

// BuscarPorID busca um usuário pelo ID
func (r *usuarioRepository) BuscarPorID(id uint64) (*models.Usuario, error) {
	var usuario models.Usuario
	if err := r.db.Select("id, nome, email, criado_em").First(&usuario, id).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}

// BuscarSenha busca a senha de um usuário pelo ID
func (r *usuarioRepository) BuscarSenha(id uint64) (string, error) {
	var usuario models.Usuario
	if err := r.db.Select("senha").First(&usuario, id).Error; err != nil {
		return "", err
	}

	return usuario.Senha, nil
}

// AtualizarUsuario atualiza os dados de um usuário existente
func (r *usuarioRepository) AtualizarUsuario(usuario *models.Usuario) error {

	updates := map[string]interface{}{
		"nome":  usuario.Nome,
		"email": usuario.Email,
	}

	if usuario.ID == 0 {
		return gorm.ErrRecordNotFound
	}
	if err := r.db.Model(&usuario).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

// DeletarUsuario deleta um usuário pelo ID
func (r *usuarioRepository) DeletarUsuario(id uint64) error {
	if id == 0 {
		return gorm.ErrRecordNotFound
	}

	if err := r.db.Delete(&models.Usuario{}, id).Error; err != nil {
		return err
	}

	return nil
}

// AtualizarSenha atualiza a senha de um usuário
func (r *usuarioRepository) AtualizarSenha(id uint64, senhaHash string) error {
	if id == 0 {
		return gorm.ErrRecordNotFound
	}

	updates := map[string]interface{}{
		"senha": senhaHash,
	}

	if err := r.db.Model(&models.Usuario{}).
		Where("id = ?", id).
		Updates(updates).Error; err != nil {
		return err
	}

	return nil
}
