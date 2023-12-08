package dispatch_points

import (
	"Wetao/internal/api_back/app/http/models/dispatch_points"
	"Wetao/internal/api_back/app/response"
	"github.com/gofiber/fiber/v2"
)

func GetDispatchesPoints(c *fiber.Ctx) error {
	points, err := dispatch_points.GetDispatchPoints()
	if err != nil {
		return response.ErrorApi(c, "Пункты выгрузки, не доступны!")
	}
	return response.SuccessApi(c, points)
}
