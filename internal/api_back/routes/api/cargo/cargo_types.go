package cargo

import (
	"Wetao/internal/api_back/app/http/handlers/cargo"
	"github.com/gofiber/fiber/v2"
)

// ApiCargoGroup /api/v1/cargo/::
func ApiCargoGroup(cargoGr fiber.Router) {
	// Get All Category and category-types
	cargoCategoryGr := cargoGr.Group("cat")
	cargoCategoryGr.Get("get", cargo.GetCargoCategoryAndTypes)
	//
	cargoPkgGr := cargoGr.Group("pkg")
	cargoPkgTypesGr := cargoPkgGr.Group("types")
	cargoPkgTypesGr.Get("get", cargo.GetCargoPkgTypes)

	wghSizesGr := cargoGr.Group("wgh_size")
	wghSizesGr.Get("grouped/get", cargo.GetCargoWghSizes)
}
