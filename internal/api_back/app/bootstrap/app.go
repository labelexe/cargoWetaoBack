package bootstrap

import (
	"Wetao/internal/api_back/app/bootstrap/providers"
	"github.com/gofiber/fiber/v2"
)

// HttpServiceProvider APP SERVICE PROVIDER
func HttpServiceProvider(app *fiber.App) {
	providers.BootAppService(app)
	providers.BootRouterService(app)
	providers.BootEventService(app)
}
