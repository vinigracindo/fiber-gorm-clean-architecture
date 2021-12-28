package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/internal/usecases/user"
)

func NewUserHandler(app fiber.Router, service user.UserUseCase) {
	app.Post("/", createUser(service))
	app.Get("/", listUsers(service))
	app.Get("/:userId", getUser(service))
}

func getUser(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("userId")
		user, err := service.GetUser(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Found",
			"data":    user,
		})
	}
}

func listUsers(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.ListUsers()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Users Found",
			"data":    users,
		})
	}
}

func createUser(service user.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		type RegisterDTO struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var registerDTO RegisterDTO
		err := c.BodyParser(&registerDTO)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		err = service.CreateUser(registerDTO.Username, registerDTO.Password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"error":  nil,
		})
	}
}
