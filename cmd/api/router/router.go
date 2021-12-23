package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	userGORMRouter "github.com/vinigracindo/fiber-gorm-clean-architecture/cmd/api/router/gorm/user"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello. I'm alive."))
	})

	api := app.Group("/api", logger.New())

	// Users
	userGORMRouter.SetupRoutes(api)
}
