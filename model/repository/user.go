package repository

import(
	"github.com/yoshiyoshiharu/go-api-server/model/entity"
)

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	GetUserByID(id string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type userRepository struct {}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUsers() ([]entity.User, error) {
	return nil, nil
}

func (r *userRepository) GetUserByID(id string) (entity.User, error) {
	return entity.User{}, nil
}

func (r *userRepository) CreateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (r *userRepository) UpdateUser(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}

func (r *userRepository) DeleteUser(id string) error {
	return nil
}
