package order

import (
	"Wetao/internal/api_back/app/bootstrap/observer"
	orderModel "Wetao/internal/api_back/app/http/models/order"
	"Wetao/internal/api_back/app/http/models/user"
	"Wetao/internal/api_back/app/response"
	"Wetao/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

type RequestAcceptBody struct {
	ProductLinkUrl string `json:"product_link_url"`
	CategorySlug   string `json:"category_slug"`
	PkgSlug        string `json:"pkg_slug"`
	Weight         string `json:"weight"`
}

func NewRequestAccept(ctx *fiber.Ctx) error {
	tgInitData := ctx.Locals("user_init_data")
	userData := tgInitData.(user.Users)

	telegramID := userData.TelegramID

	reqData := &RequestAcceptBody{}
	if ctx.BodyParser(reqData) != nil {
		return response.ErrorApi(ctx, "Запрос не доступен!")
	}
	order := orderModel.Orders{
		Idx:              pkg.GenerateUniqueID(),
		UserID:           telegramID,
		CatID:            1,
		PkgID:            1,
		TypeBuyback:      "we_buy",
		TotalAmountPrice: 1500,
		ProductLinkURL:   reqData.ProductLinkUrl,
		//
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "pending_accepted", // Начальный статус заявки
	}

	parts := strings.Split(reqData.Weight, ",")
	price, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		log.Err(err)
	}
	order.WeightSizes = orderModel.WeightSizes{
		Name:  parts[1],
		Price: price,
	}

	//
	order.Slug = orderModel.GenerateUniqueSlug("we_buy")

	createOrderData, err := orderModel.InsertNewRequestByOrder(&order)
	if err != nil {
		log.Err(err)
		return response.ErrorApi(ctx, "Запрос не создан!")
	}

	//send new event
	observer.NotifyObserver("app.order.new_request", createOrderData)

	return response.SuccessApi(ctx, fiber.Map{
		"data": createOrderData,
	})
}
