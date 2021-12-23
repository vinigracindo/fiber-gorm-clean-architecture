package repository

import (
	"time"

	"github.com/vinigracindo/fiber-gorm-clean-architecture/entity"
	usecaseuser "github.com/vinigracindo/fiber-gorm-clean-architecture/usecase/user"
	"gorm.io/gorm"
)

type UserGORM struct {
	gorm.Model
	ID        entity.ID `gorm:"type:uuid"`
	Username  string
	Password  string
	CreatedAt time.Time
}

// Set tablename (GORM)
func (UserGORM) TableName() string {
	return "users"
}

func (u *UserGORM) setUserGORM(entityUser *entity.User) {
	u.ID = entityUser.ID
	u.Username = entityUser.Username
	u.Password = entityUser.Password
	u.CreatedAt = entityUser.CreatedAt
}

func (u UserGORM) getEntityUser() *entity.User {
	return &entity.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) usecaseuser.UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(entityUser *entity.User) error {
	u := &UserGORM{}
	u.setUserGORM(entityUser)

	err := r.DB.Create(&u).Error
	if err != nil {
		return err
	}

	return nil
}
