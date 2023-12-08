package payment

import (
	currencyModel "Wetao/internal/api_back/app/http/models/currency"
	"Wetao/pkg/cache"
	"Wetao/pkg/currency"
	"github.com/gofiber/fiber/v2"
	"time"
)

type CurrencyResponse struct {
	Success  bool        `json:"success"`
	ErrorMsg string      `json:"error_msg"`
	Result   interface{} `json:"result"`
}

func getCacheByCurrency() ([]currencyModel.CurRate, error) {
	cache, err := cache.RememberCache("currency_rates", 45*time.Minute, func() (interface{}, error) {
		_, err := currency.UpdateRtByCurAndGet([]string{})
		if err != nil {
			return nil, err
		}

		manyCurrencyDirRates, _, err := currencyModel.GetManyCurrencyDirRates([]string{"USD", "CNY"})
		if err != nil {
			return []currencyModel.CurRate{}, err
		}
		//
		return manyCurrencyDirRates, nil
	})
	if err != nil {
		return nil, err
	}

	return cache.([]currencyModel.CurRate), nil
}

// GetCurrencyInfo Получение информации о курсах
func GetCurrencyInfo(ctx *fiber.Ctx) error {
	currencyInfoCache, err := getCacheByCurrency()
	if err != nil {
		ctx.JSON(CurrencyResponse{
			Success:  false,
			ErrorMsg: "Not get currency!",
		})
	}
	//
	ctx.JSON(CurrencyResponse{
		Success:  true,
		ErrorMsg: "",
		Result:   currencyInfoCache,
	})

	return nil
}
