package jwt_auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	modelUser "retryh-billing-go/app/models/user"
	"time"
)

type userCredential struct {
	Username []byte `json:"username"`
	Password []byte `json:"password"`
	jwt.RegisteredClaims
}

var (
	SecretToken string
)

//const SECRET_TOKEN = "kdjfdjklejrkrej"

func GenerateNewPairs() {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("GenerateKey Pair: %v", err)
	}
	//
	SecretToken = string(key)
}

func GenerateJwtTokenWithClaims(username string, password []byte) (string, jwt.Claims, error) {
	// Create the Claims
	//claims := jwt.MapClaims{
	//	"name":  username,
	//	"admin": true,
	//	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	//}

	expireAt := time.Now().Add(36 * time.Hour)

	claims := &userCredential{
		Username: []byte(username),
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SecretToken))
	return t, claims.RegisteredClaims, err
}

func VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretToken), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// MiddlewareJwtAuthRoute - Middleware проверки авторизации
// Пишем его перед маршрутами-группы методов или после метода [GET,POST...]
// С Указанием на его него ссылки
func MiddlewareJwtAuthRoute(app *fiber.App, router fiber.Router) *fiber.App {
	// JWT Middleware
	router.Use(jwtware.New(
		jwtware.Config{
			//Обработка успешной авторизации
			SuccessHandler: func(c *fiber.Ctx) error {
				//fmt.Println("Auth->Success->Handle-->", c)
				log.Println("user", "claims:", c)
				return c.Next()
			},
			//Обработка ошибок
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Not Auth!",
				})
			},
			SigningKey: jwtware.SigningKey{
				JWTAlg: jwtware.HS512,
				Key:    []byte(SecretToken),
			},
			//TokenLookup: "header:ex",
			//TokenLookup: "header:Authorization",
		},
	))

	return app

}

//Tool

// GetClaimByUser - Получения сессионных полей пользователя
func GetClaimByUser(c *fiber.Ctx, field string) interface{} {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims[field]
}

func GetUserDataByClaim(c *fiber.Ctx) (modelUser.Users, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	usernameBase64 := claims["username"].(string)
	username, err := base64.StdEncoding.DecodeString(usernameBase64)
	//
	byUsername, err := modelUser.GetUserByUsername(modelUser.AuthUser{Username: string(username)})
	if err != nil {
		return modelUser.Users{}, err
	} else {
		return byUsername, nil
	}
}
