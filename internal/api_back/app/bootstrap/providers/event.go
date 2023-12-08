package providers

import (
	"Wetao/internal/api_back/app/bootstrap/observer"
	"Wetao/internal/api_back/app/event"
	eventOrderRequest "Wetao/internal/api_back/app/event/orders/request"

	"github.com/gofiber/fiber/v2"
)

func BootEventService(app *fiber.App) {
	observer.RegisterObserver("app.user.register_user", &event.RegUserListener{})
	observer.RegisterObserver("app.user.new_customer", &event.NewCustomerListener{})
	observer.RegisterObserver("app.order.new_request", &eventOrderRequest.NewRequestOrderListener{})
}
