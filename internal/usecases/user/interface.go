package user

import "github.com/vinigracindo/fiber-gorm-clean-architecture/internal/entities"

type Reader interface {
	Get(userID string) (*entities.User, error)
	GetByUsername(username string) (*entities.User, error)
	List() ([]*entities.User, error)
}

type Writer interface {
	Create(u *entities.User) error
	//Update(u *entities.User) error
	//Delete(u *entities.User) error
}

type UserRepository interface {
	Reader
	Writer
}

type UserUseCase interface {
	CreateUser(username, password string) error
	ListUsers() ([]*entities.User, error)
	GetUser(id string) (*entities.User, error)
}
