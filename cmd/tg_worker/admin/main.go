package admin

import (
	"Wetao/pkg"
	"fmt"
	tg_bot_api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"os"
)

func TgBotInit() *tg_bot_api.BotAPI {
	//Env
	pkg.EnvLoadInit()

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Err(fmt.Errorf("TELEGRAM_BOT_TOKEN not set in .env file"))
	}

	//botToken := "6473213447:AAExR_cZYHQGhrWTT0oxGBcGuPiCnZxeHxY" // Replace with your bot's API token
	bot, err := tg_bot_api.NewBotAPI(botToken)
	if err != nil {
		log.Err(err)
	}

	bot.Debug = true // Enable debugging (optional)

	log.Printf("Authorized as @%s", bot.Self.UserName)

	return bot
}
