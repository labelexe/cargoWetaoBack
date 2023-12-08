package response

import "github.com/gofiber/fiber/v2"

type CurrencyResponse struct {
	Success  bool        `json:"success"`
	ErrorMsg string      `json:"error_msg"`
	Result   interface{} `json:"result"`
}

func SuccessApi(ctx *fiber.Ctx, resultData interface{}) error {
	return ctx.JSON(CurrencyResponse{
		Success:  true,
		ErrorMsg: "",
		Result:   resultData,
	})
}

func ErrorApi(ctx *fiber.Ctx, errorMsg string) error {
	return ctx.JSON(CurrencyResponse{
		Success:  false,
		ErrorMsg: errorMsg,
		Result:   nil,
	})
}
