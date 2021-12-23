package userUsecase

import "github.com/vinigracindo/fiber-gorm-clean-architecture/entity"

type Reader interface {
}

type Writer interface {
	CreateUser(u *entity.User) error
}

type UserRepository interface {
	Reader
	Writer
}

type UserUseCase interface {
	CreateUser(u *entity.User) error
}
