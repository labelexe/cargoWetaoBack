package orders

import (
	"Wetao/internal/api_back/app/http/handlers/order"
	"github.com/gofiber/fiber/v2"
)

// ApiOrdersGroup /api/v1/orders/::
func ApiOrdersGroup(ordersGr fiber.Router) {

	// Checkout Cart
	ordersGr.Post("/checkout", order.CheckoutOrderByUser)

	// Request
	ApiOrdersRequestGroup(ordersGr)
}
