package orders

import (
	"Wetao/internal/api_back/app/http/handlers/order"
	"github.com/gofiber/fiber/v2"
)

// ApiOrdersRequestGroup /api/v1/orders/requests/::
func ApiOrdersRequestGroup(ordersGr fiber.Router) {
	orderRequestsGr := ordersGr.Group("requests")
	orderRequestsGr.Post("accept", order.NewRequestAccept)
}
