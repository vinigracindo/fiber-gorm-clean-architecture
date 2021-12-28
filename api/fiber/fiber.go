package fiber

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/api/fiber/internal/handler"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/infrastructure/gorm/database"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/infrastructure/gorm/repository"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/usecases/user"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/pkg/config"
	"gorm.io/driver/postgres"
)

func Run(port int) {
	config.LoadEnvironmentFile(".env")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBUser,
		config.Env.DBPass,
		config.Env.DBName,
	)

	gormdb, err := database.ConnectGORMDB(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Database connection error: $s", err)
	}

	// Creates a new Fiber instance.
	app := fiber.New(fiber.Config{
		AppName:      "Fiber Gorm Clean Architecture",
		ServerHeader: "Fiber",
	})

	// Use global middlewares.
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Create repositories.
	userRepository := repository.NewUserGORMRepository(gormdb)

	// Create all of our services.
	userService := user.NewService(userRepository)

	api := app.Group("/api")

	// Prepare our endpoints for the API.
	handler.NewUserHandler(api.Group("/v1/users"), userService)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 3000.
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
