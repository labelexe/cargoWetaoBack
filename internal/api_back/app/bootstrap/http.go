package bootstrap

import "github.com/gofiber/fiber/v2"

func HttpAppServer() *fiber.App {
	app := fiber.New(fiber.Config{
		// Style Server
		AppName:      "WetaoDelSvc",
		ServerHeader: "WetaoDelSvc",
		// Config Server
		CaseSensitive: true,
		StrictRouting: true,
	})
	return app
}
