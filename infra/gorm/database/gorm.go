package database

import (
	"fmt"
	"log"

	"github.com/vinigracindo/fiber-gorm-clean-architecture/config"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/infra/gorm/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm.DB instance
var GORMDB *gorm.DB

func ConnectGORMDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBUser,
		config.Env.DBPass,
		config.Env.DBName,
	)
	GORMDB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("failed to connect database")
	}
	GORMDB.AutoMigrate(&repository.UserGORM{})
}
