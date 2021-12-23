package userHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinigracindo/fiber-gorm-clean-architecture/entity"
	usecaseuser "github.com/vinigracindo/fiber-gorm-clean-architecture/usecase/user"
)

func CreateUser(service usecaseuser.UserUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entity.User
		err := c.BodyParser(&requestBody)

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		user, err := entity.NewUser(requestBody.Username, requestBody.Password)

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		dberr := service.CreateUser(user)
		return c.JSON(&fiber.Map{
			"status": "ok",
			"error":  dberr,
		})
	}
}
