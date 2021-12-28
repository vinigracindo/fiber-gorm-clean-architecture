package user

import (
	"errors"

	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/entities"
)

type service struct {
	userRepository UserRepository
}

func NewService(r UserRepository) UserUseCase {
	return &service{
		userRepository: r,
	}
}

func (s *service) CreateUser(username, password string) error {
	_, err := s.userRepository.GetByUsername(username)
	if err == nil {
		return errors.New("User already exists")
	}
	u, err := entities.NewUser(username, password)
	if err != nil {
		return err
	}
	err = s.userRepository.Create(u)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ListUsers() ([]*entities.User, error) {
	return s.userRepository.List()
}

func (s *service) GetUser(id string) (*entities.User, error) {
	return s.userRepository.Get(id)
}
