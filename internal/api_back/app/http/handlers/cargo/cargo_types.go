package cargo

import (
	"Wetao/internal/api_back/app/http/models/cargo"
	"Wetao/internal/api_back/app/response"
	"github.com/gofiber/fiber/v2"
)

func GetCargoCategoryAndTypes(ctx *fiber.Ctx) error {
	cc, err := cargo.GetAllCargoCategoriesAndTypes()
	if err != nil {
		return response.ErrorApi(ctx, "Категории не доступны!")
	}
	return response.SuccessApi(ctx, cc)
}

func GetCargoPkgTypes(ctx *fiber.Ctx) error {
	cc, err := cargo.GetAllCargoPkgTypes()
	if err != nil {
		return response.ErrorApi(ctx, "Категории не доступны!")
	}
	return response.SuccessApi(ctx, cc)
}
