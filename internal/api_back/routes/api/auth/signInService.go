package auth

//
//import (
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	"golang.org/x/crypto/bcrypt"
//	userModel "retryh-billing-go/app/models/user"
//	"retryh-billing-go/app/service/api/auth/jwt_auth"
//)
//
//type UserSignIn struct {
//	User     string `json:"login_email"`
//	Password string `json:"password"`
//}
//
//// SignIn @Router /api/v1/auth/signin
//func SignIn(c *fiber.Ctx) error {
//	//Тело-запроса
//	userReqs := new(UserSignIn)
//	if err := c.BodyParser(userReqs); err != nil {
//		fmt.Println("error = ", err)
//		return c.SendStatus(200)
//	}
//
//	//createUser, err := userModel.CreateUser(userReqs.User, userReqs.Password)
//	//if err != nil {
//	//	return err
//	//}
//
//	// Получение пользователя
//	var user userModel.Users
//
//	user, err := userModel.GetUserByUsername(userModel.AuthUser{Username: userReqs.User})
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//			"error": "Invalid credentials #592",
//		})
//	}
//
//	// Сверка шифра пароля
//	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReqs.Password))
//	if err != nil {
//		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
//			"error": "Invalid credentials #842",
//		})
//	}
//
//	// Выдача подписанного токена
//	t, claims, err := jwt_auth.GenerateJwtTokenWithClaims(user.Username, []byte(userReqs.Password))
//	if err != nil {
//		fmt.Println("Error Jwt Auth | User {}", user)
//	}
//
//	user, err = userModel.GetUserByUsername(userModel.AuthUser{Username: user.Username})
//	userAndUsd := user.CalcExchangesRates("USD")
//	//
//	if err != nil {
//		return c.JSON(fiber.Map{"info": claims, "user": fiber.Map{}, "auth": "Bearer", "act": t})
//	} else {
//		return c.JSON(fiber.Map{"info": claims, "user": userAndUsd, "auth": "Bearer", "act": t})
//	}
//}
//
//func SignInView(c *fiber.Ctx) error {
//	return c.Render(
//		"auth/signin", fiber.Map{
//			"Title": "Авторизация",
//		},
//		"layout/main")
//}
