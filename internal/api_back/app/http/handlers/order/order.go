package order

import (
	"Wetao/internal/api_back/app/bootstrap/observer"
	"github.com/gofiber/fiber/v2"
)

// CheckoutOrderByUser
// TODO: CHECK - (
// TODO: > COUNT PRODUCT QUANTITY - /
// TODO: > BALANCE(Bonus) -
// TODO: >
// TODO: )
func CheckoutOrderByUser(ctx *fiber.Ctx) error {
	observer.NotifyObserver("app.user.register_user", fiber.Map{"пример данных": "data"})
	return nil
}
