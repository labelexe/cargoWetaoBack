package routes

import (
	"Wetao/internal/api_back/app/http/handlers/auth"
	"Wetao/internal/api_back/app/http/middleware/webapp/validate"
	"Wetao/internal/api_back/routes/api/cargo"
	dispatch_points "Wetao/internal/api_back/routes/api/dispatch_points"
	orderRouter "Wetao/internal/api_back/routes/api/orders"
	"Wetao/internal/api_back/routes/api/payment"
	productRouter "Wetao/internal/api_back/routes/api/product"
	"github.com/gofiber/fiber/v2"
)

// TODO: https://docs.telegram-mini-apps.com/packages/golang/init-data-golang

// ApiRouterGroupInit Инициализация api роутера
func ApiRouterGroupInit(app *fiber.App, apiGroup fiber.Router) {
	//First Version
	apiV1Gr := apiGroup.Group("v1")
	//Routes
	ApiV1GroupRoutes(apiV1Gr)
}

func ApiV1GroupRoutes(apiFVGr fiber.Router) {
	apiFVGr.Use(validate.NewTwaAuthMiddleware)

	authGr := apiFVGr.Group("auth")
	authGr.Post("/profile", auth.ProfileHandler)
	//
	//authGr.Post("/tg/sign", func(c *fiber.Ctx) error {
	//	authHeaderInitData := c.Get("sessionT")
	//	tg, err := tg_pkg.CheckAuthorization(authHeaderInitData)
	//	if err != nil {
	//		//return err
	//		c.JSON(fiber.Map{
	//			"err_msg": err,
	//			"err":     true,
	//		})
	//	}
	//	c.JSON(fiber.Map{
	//		"err_msg": "",
	//		"err":     false,
	//		"sg":      tg,
	//	})
	//	//fmt.Println(tg)
	//	//data, _ := url.ParseQuery(authHeaderInitData)
	//	return nil
	//})

	//
	productGr := apiFVGr.Group("product")
	productRouter.ApiProductGroup(productGr)
	//
	orderGr := apiFVGr.Group("orders")
	orderRouter.ApiOrdersGroup(orderGr)

	// Payment / Currency rate data!
	paymentGr := apiFVGr.Group("payment")
	payment.ApiPaymentGroup(paymentGr)

	cargoGr := apiFVGr.Group("cargo")
	cargo.ApiCargoGroup(cargoGr)

	dispatchPointGr := apiFVGr.Group("dispatch_p")
	dispatch_points.ApiDispatchPointsGroup(dispatchPointGr)
	//
}
