package repository

import(
	"github.com/yoshiyoshiharu/go-api-server/model/entity"
)

type UserRepository interface {
	GetUsers() ([]entity.UserEntity, error)
	GetUserByID(id string) (entity.UserEntity, error)
	CreateUser(user entity.UserEntity) (entity.UserEntity, error)
	UpdateUser(user entity.UserEntity) (entity.UserEntity, error)
	DeleteUser(id string) error
}

type userRepository struct {}

// UserRepositoryのインタフェースを返す
// 利用者は5つのメソッドを実装した構造体が返ることのみを知っていればよい
func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUsers() ([]entity.UserEntity, error) {
	return nil, nil
}

func (r *userRepository) GetUserByID(id string) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}

func (r *userRepository) CreateUser(user entity.UserEntity) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}

func (r *userRepository) UpdateUser(user entity.UserEntity) (entity.UserEntity, error) {
	return entity.UserEntity{}, nil
}

func (r *userRepository) DeleteUser(id string) error {
	return nil
}
