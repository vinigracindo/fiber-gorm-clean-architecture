package database

import (
	"fmt"

	"github.com/vinigracindo/fiber-gorm-clean-architecture/config"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/infra/gorm/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB gorm connector
var GORMDB *gorm.DB

func ConnectGORMDB() {
	var err error
	host := config.Config("DB_HOST")
	port := config.Config("DB_PORT")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	dbName := config.Config("DB_NAME")

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	// Connect to the DB and initialize the DB variable
	GORMDB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	GORMDB.AutoMigrate(&repository.UserGORM{})
}
