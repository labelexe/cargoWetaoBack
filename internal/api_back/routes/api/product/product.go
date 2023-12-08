package product

import (
	"Wetao/internal/api_back/app/http/handlers/product"
	"github.com/gofiber/fiber/v2"
)

// ApiProductGroup /api/v1/product/::
func ApiProductGroup(productGr fiber.Router) {

	// Sort Product By Effect
	productGr.Get("sortByEff", product.SortByEffect)

	// Recommendation Products
	productGr.Get("recommend", product.RecommendProduct)

	// All Products
	productGr.Get("all", product.AllProduct)

	// Get Single Product
	productGr.Get("show/:id", product.ShowProduct)
}
