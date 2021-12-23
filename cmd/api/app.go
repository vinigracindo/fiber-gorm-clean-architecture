package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/cmd/api/router"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/infra/gorm/database"
)

func Run(port int) {
	database.ConnectGORMDB()
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Ping! Pong!"))
	})
	router.SetupRoutes(app)
	app.Listen(fmt.Sprintf(":%d", port))
}
