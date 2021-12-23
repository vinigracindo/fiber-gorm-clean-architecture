package userGORMRouter

import (
	"github.com/gofiber/fiber/v2"
	userHandler "github.com/vinigracindo/fiber-gorm-clean-architecture/cmd/api/handler/user"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/infra/gorm/database"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/infra/gorm/repository"
	userUsecase "github.com/vinigracindo/fiber-gorm-clean-architecture/usecase/user"
)

func routes(app fiber.Router, service userUsecase.UserUseCase) {
	user := app.Group("/v1/users")
	user.Post("/", userHandler.CreateUser(service))
}

func SetupRoutes(app fiber.Router) {
	userRepo := repository.NewUserRepository(database.GORMDB)
	userService := userUsecase.NewService(userRepo)
	routes(app, userService)
}
