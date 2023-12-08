package auth

import (
	"Wetao/internal/api_back/app/http/models/user"
	"Wetao/internal/api_back/app/response"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	return nil
}

func ProfileHandler(c *fiber.Ctx) error {
	tgInitData := c.Locals("user_init_data")
	return response.SuccessApi(c, tgInitData.(user.Users))
}
