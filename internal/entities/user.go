package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/pkg/id"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        id.ID     `valid:"required" json:"id"`
	Username  string    `valid:"required" json:"username"`
	Password  string    `valid:"required" json:"password"`
	CreatedAt time.Time `valid:"required" json:"created_at"`
}

func (u User) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func newPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func NewUser(username string, password string) (*User, error) {
	pwd, err := newPasswordHash(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:        id.NewID(),
		Username:  username,
		Password:  pwd,
		CreatedAt: time.Now(),
	}
	_, err = user.IsValid()
	if err != nil {
		return nil, err
	}

	return user, nil
}
