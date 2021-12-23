package userUsecase

import (
	"github.com/vinigracindo/fiber-gorm-clean-architecture/entity"
)

type service struct {
	userRepository UserRepository
}

func NewService(r UserRepository) UserUseCase {
	return &service{
		userRepository: r,
	}
}

func (s *service) CreateUser(user *entity.User) error {
	return s.userRepository.CreateUser(user)
}
