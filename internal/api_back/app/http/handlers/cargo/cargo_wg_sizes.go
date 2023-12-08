package cargo

import (
	"Wetao/internal/api_back/app/http/models/cargo"
	"Wetao/internal/api_back/app/response"
	"github.com/gofiber/fiber/v2"
)

func GetCargoWghSizes(ctx *fiber.Ctx) error {
	cc, err := cargo.GetAllWghSizesDistName()
	if err != nil {
		return response.ErrorApi(ctx, "Размеры не доступны!")
	}
	return response.SuccessApi(ctx, cc)
}
