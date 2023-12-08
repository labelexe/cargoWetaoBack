package request

import (
	"Wetao/cmd/tg_worker/admin"
	orderModel "Wetao/internal/api_back/app/http/models/order"
	"fmt"
	tg_bot_api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"strings"
)

//5159648991:AAGoNNIAaCH47DQqjVDagt1kqnMRnPJR4GM

func TelegramSendMsgNewOrderReq(item orderModel.Orders) {
	chatID := 1248134771
	//
	var arrMsg []string
	//var tags []string
	//var paramTagsText = ite

	var typeBuyBackStr string

	switch item.TypeBuyback {
	case "we_buy":
		typeBuyBackStr = "Выкупаем мы! +5%"
	case "you_buy":
		typeBuyBackStr = "Выкупает покупатель!"
	}

	// Добавляем строки в срез
	arrMsg = append(arrMsg,
		"🔥📦 Новая заявка на карго:\n",
		"----", "📌 Кто выкупает: <b>", typeBuyBackStr, "</b>----\n")

	//if len(paramTagsText) > 0 {
	//	if paramTagsText != "" {
	//		tags = strings.Split(paramTagsText, ",")
	//	}
	//
	//	arrMsg = append(arrMsg, fmt.Sprintf("👀	Теги: %s \n", "#"+strings.Join(tags, ",#")))
	//}

	//arrMsg = append(arrMsg, fmt.Sprintf("🕓 Дата публикации: %s \n", item.Date), "----\n")
	arrMsg = append(arrMsg, fmt.Sprintf("*** <b>Номер заказа:</b> #idx_%s", item.Idx), "----\n")
	//
	arrMsg = append(arrMsg, fmt.Sprintf("🛒 Ссылка на товар(объект) — %s", item.ProductLinkURL), "----")
	arrMsg = append(arrMsg, fmt.Sprintf("🛒 Объём груза — %s %s ", item.WeightSizes.Name, "KG"), "----")
	arrMsg = append(arrMsg, fmt.Sprintf("✈️ Отправка-Получение — %s", "<b>Gaunho — Moscow</b>"), "----")
	arrMsg = append(arrMsg, fmt.Sprintf("🥸 Заказчи — ID #usr_%d ", item.UserID), "----")
	arrMsg = append(arrMsg, fmt.Sprintf("🎀 Тип упаковки — %d ", item.PkgID), "----")
	arrMsg = append(arrMsg, fmt.Sprintf("✈️ Категория товара — %d", item.CatID), "----")
	arrMsg = append(arrMsg, fmt.Sprintf("🧳 Текущий статус #%s ", item.Status), "----")
	arrMsg = append(arrMsg, fmt.Sprintf("💸 Общая сумма — %v", item.TotalAmountPrice), "----")
	//
	arrMsg = append(arrMsg, fmt.Sprintf("TAGS: #%v, #idx_%v, #usr_%d", item.TypeBuyback, item.Idx, item.UserID), "----")
	//arrMsg = append(arrMsg, fmt.Sprintf("🕓 Дата публикации: %s \n", item.Date), "----\n")

	//if len(item.Description) >= 480 {
	//	arrMsg = append(arrMsg, fmt.Sprintf("%s \n", item.Description[0:420]))
	//} else {
	//	arrMsg = append(arrMsg, fmt.Sprintf("%s \n", item.Description))
	//}

	// Объединяем все строки в одну строку с переносами строк
	combinedMsgString := strings.Join(arrMsg, "\n")

	// Create a message configuration
	msg := tg_bot_api.NewMessage(int64(chatID), combinedMsgString)
	msg.ParseMode = tg_bot_api.ModeHTML

	var keyboardRow []tg_bot_api.InlineKeyboardButton
	var keyboardRow2 []tg_bot_api.InlineKeyboardButton
	var btnUrl = item.ProductLinkURL

	keyboardRow = tg_bot_api.NewInlineKeyboardRow(
		tg_bot_api.NewInlineKeyboardButtonURL(fmt.Sprintf("🛍 Просмотр товара \n"), item.ProductLinkURL),
	)
	keyboardRow2 = tg_bot_api.NewInlineKeyboardRow(
		tg_bot_api.NewInlineKeyboardButtonURL(fmt.Sprintf("🤝 Принять заявку - ц. %f Руб.", item.TotalAmountPrice), btnUrl),
	)
	//
	msg.ReplyMarkup = tg_bot_api.NewInlineKeyboardMarkup(keyboardRow, keyboardRow2)
	//msg.ReplyMarkup = tg_bot_api.NewInlineKeyboardMarkup(keyboardRow)

	// Send the text message
	messageSend, err := admin.
		TgBotInit().
		Send(msg)
	if err != nil {
		log.Err(fmt.Errorf("err tg: %v", err))
	} else {
		log.Printf("tg_msg_id:", messageSend.MessageID)
	}

}
