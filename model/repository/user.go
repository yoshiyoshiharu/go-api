package repository

import (
	"fmt"
	"log"
	"errors"
	"database/sql"

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
	users := []entity.UserEntity{}

	rows, err := Db.Query("SELECT * FROM users")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	for rows.Next() {
		user := entity.UserEntity{}
		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
		if err != nil {
			log.Print(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(id string) (entity.UserEntity, error) {
	user := entity.UserEntity{}

  row := Db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("User not found")
			return entity.UserEntity{}, err
		} else {
			log.Print(err)
			return entity.UserEntity{}, err
		}
	}

	return user, nil
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
