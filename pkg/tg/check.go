package tg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"os"
	"time"
)

func subtleCompare(a, b string) int {
	// Constant time comparison to avoid timing attacks
	if len(a) != len(b) {
		return 1
	}

	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}

	return int(result)
}

func CheckTelegramAuthorization(authHeaderInitData string) (bool, error) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Err(fmt.Errorf("TELEGRAM_BOT_TOKEN not set in .env file"))
	}

	initData, err := initdata.Parse(authHeaderInitData)
	if err != nil {
		return false, err
	}

	secretKey := sha256.Sum256([]byte(botToken))
	hash := hmac.New(sha256.New, secretKey[:])
	hash.Write([]byte(authHeaderInitData))
	calculatedHash := hex.EncodeToString(hash.Sum(nil))

	if subtleCompare(calculatedHash, initData.Hash) != 0 {
		return false, fmt.Errorf("data is NOT from Telegram")
	}

	authDate := initData.AuthDateRaw
	if authDate != 0 || time.Now().Unix()-int64(authDate) > 86400 {
		return false, fmt.Errorf("data is outdated")
	}

	return true, nil
}

func SignAuthTg(initData string) (string, error) {
	// Signing timestamp.
	authDate := time.Now()
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Err(fmt.Errorf("TELEGRAM_BOT_TOKEN not set in .env file"))
	}
	//fmt.Println("SignQueryString: ")
	return initdata.SignQueryString(initData, botToken, authDate)
}

func CheckAuthorization(authHeaderInitData string) (initdata.InitData, error) {
	signAuthTg, err := SignAuthTg(authHeaderInitData)
	if err != nil {
		return initdata.InitData{}, err
	}
	fmt.Println(signAuthTg)

	return initdata.Parse(authHeaderInitData)
}

func VerifyAuthData(initData string, authDate string, hash string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Err(fmt.Errorf("TELEGRAM_BOT_TOKEN not set in .env file"))
	}

	// Подготовка строки для вычисления хеша
	authString := fmt.Sprintf("%d:%s", authDate, initData)

	// Вычисление хеша
	calculatedHash := hmac.New(sha256.New, []byte(botToken))
	calculatedHash.Write([]byte(authString))
	expectedHash := hex.EncodeToString(calculatedHash.Sum(nil))

	// Сравнение хешей
	if expectedHash != hash {
		return errors.New("invalid authorization data")
		//return c.Status(fiber.StatusUnauthorized).SendString("Invalid authorization data")
	}
	return nil
}
