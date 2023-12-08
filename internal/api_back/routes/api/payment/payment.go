package payment

import (
	"Wetao/internal/api_back/app/http/handlers/payment"
	"github.com/gofiber/fiber/v2"
)

// ApiPaymentGroup /api/v1/payment/::
func ApiPaymentGroup(paymentGr fiber.Router) {
	//PaymentForm End.

	//Currency Exc data End.
	excRate := paymentGr.Group("currency")
	excRate.Get("rate", payment.GetCurrencyInfo)
}
