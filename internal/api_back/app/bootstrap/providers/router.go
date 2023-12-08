package providers

import (
	"Wetao/internal/api_back/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func BootRouterService(app *fiber.App) {
	log.Print("Http Init Router.. Ok")
	//
	apiGroup := app.Group("api")
	routes.ApiRouterGroupInit(app, apiGroup)
}
