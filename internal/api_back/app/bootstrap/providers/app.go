package providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"
)

func BootAppService(app *fiber.App) {
	log.Print("App Service Provider .. Ok")

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", //TODO: Close Origins
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}
