package database

import (
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/infrastructure/gorm/repository"
	"gorm.io/gorm"
)

func ConnectGORMDB(dialector gorm.Dialector) (*gorm.DB, error) {
	var gormDB *gorm.DB
	gormDB, err := gorm.Open(dialector)
	if err != nil {
		return nil, err
	}
	gormDB.AutoMigrate(&repository.UserGORM{})
	return gormDB, nil
}
