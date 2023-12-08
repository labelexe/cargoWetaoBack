package validate

import (
	"Wetao/internal/api_back/app/http/models/order"
	userModel "Wetao/internal/api_back/app/http/models/user"
	"Wetao/internal/database"
	"Wetao/pkg/tg"
	"fmt"
	"github.com/gofiber/fiber/v2"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/url"
	"os"
	"time"
)

const WEB_APP_DATA_CONST = "WebAppTest"
const TELEGRAM_BOT_TOKEN = "5159648991:AAGoNNIAaCH47DQqjVDagt1kqnMRnPJR4GM"

var SKIP_VALIDATTION = os.Getenv("SKIP_VALIDATION") == "true"

type TgWebAppHeader struct {
	sessionT string `reqHeader:"sessionT"`
}

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Prefix defines a URL prefix added before "/debug/pprof".
	// Note that it should start with (but not end with) a slash.
	// Example: "/federated-fiber"
	//
	// Optional. Default: ""
	Prefix string
}

func validateTg(initData string) error {
	// Init data in raw format.
	// Define how long since init data generation date init data is valid.
	expIn := 24 * time.Hour
	// Will return error in case, init data is invalid. To see,
	// which error could be returned, see errors.go file.
	return initdata.Validate(initData, TELEGRAM_BOT_TOKEN, expIn)
}

func NewValidIntDataNew(c *fiber.Ctx) error {
	// Получаем значение Authorization из заголовка
	authHeaderInitData := c.Get("sessionT")
	data, _ := url.ParseQuery(authHeaderInitData)
	fmt.Println(data)
	// Проверяем, что Authorization присутствует
	if authHeaderInitData == "" {
		// Если отсутствует, возвращаем ошибку
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	//fmt.Println(authorizationHeader)

	// Telegram Bot secret key.
	initDataUser, err := initdata.Parse(authHeaderInitData)
	if err != nil {
		// Если отсутствует, возвращаем ошибку
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	sing, err := tg.SignAuthTg(authHeaderInitData)
	if err != nil {
		return err
	}

	fmt.Println("Sign Tg :", sing)

	//err = validateTg(authHeaderInitData)
	//if err != nil {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//		"error": "Not Validate",
	//	})
	//}

	fmt.Printf("[Init DATA TG] %v", initDataUser)
	//fmt.Println(authorizationHeader, initdata.Validate(authorizationHeader, token, expIn))

	// Здесь вы можете использовать значение Authorization по вашему усмотрению
	// Например, вы можете добавить его в контекст или выполнять дополнительные проверки

	// Продолжаем выполнение следующих middleware и обработчиков
	return c.Next()
}

func NewTwaAuthMiddleware(c *fiber.Ctx) error {
	// Raw init data from the request header (https://docs.twa.dev/docs/launch-params/init-data).
	initData := c.Get("X-Init-Data")

	fmt.Println(initData)

	if initData == "" {
		return c.
			Status(400).
			JSON(fiber.Map{"error": "please pass signed init data in the X-Init-Data header"})
	}

	// How long the init data is valid
	expIn := 24 * time.Hour

	if !SKIP_VALIDATTION {
		if err := initdata.Validate(initData, TELEGRAM_BOT_TOKEN, expIn); err != nil {
			return c.
				Status(400).
				JSON(fiber.Map{"error": fmt.Sprintf("[Validate] can't validate init data: %s", err.Error())})
		}
	}

	data, err := initdata.Parse(initData)
	if err != nil {
		return c.
			Status(400).
			JSON(fiber.Map{"error": fmt.Sprintf("[Parse] can't parse init data: %s", err.Error())})
	}

	db := database.GetDB()
	var users = userModel.Users{
		TelegramID: data.User.ID,
		Username:   data.User.Username,
		IsActive:   true,
		FirstName:  data.User.FirstName,
		LastName:   data.User.LastName,
		Status:     "active",
		AuthToken:  "",
		IsPremium:  data.User.IsPremium,
		PhotoURL:   data.User.PhotoURL,
		UserOrders: []order.Orders{},
		//UserOrdersInfo: struct {
		//	PendingCount   int64 `json:"pending_count"`
		//	AcceptedSendCount int64 `json:"accepted_sended"`
		//}
	}

	// Пример использования метода LoadUserOrdersInfo

	//users.UserOrdersInfo.AcceptedSendCount = 0
	//users.UserOrdersInfo.PendingCount = 0

	if err := db.Preload("UserOrders").
		Preload("UserOrders.OrderPackage").
		Preload("UserOrders.OrderCategory").
		Preload("UserOrders.OrderCategory.CargoWithType").
		FirstOrCreate(&users, users).Error; err != nil {
		return c.
			Status(500).
			JSON(fiber.Map{"error": fmt.Sprintf("[Create or First] can't upsert user: %s", err.Error())})
	}

	err = users.GetUserOrdersInfo(db)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}

	c.Locals("user_init_data", users)
	return c.Next()
}
