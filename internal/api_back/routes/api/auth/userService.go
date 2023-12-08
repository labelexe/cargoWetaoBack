package auth

//
//import (
//	"github.com/gofiber/fiber/v2"
//	"retryh-billing-go/app/service/api/auth/jwt_auth"
//)
//
//// GetProfile @Router /api/v1/auth/user
//func GetProfile(c *fiber.Ctx) error {
//	user, err := jwt_auth.GetUserDataByClaim(c)
//	userAndExc := user.CalcExchangesRates("USD")
//	if err != nil {
//		return c.JSON(fiber.Map{
//			"success": false,
//			"user":    nil,
//		})
//	} else {
//		return c.JSON(fiber.Map{
//			"success": true,
//			"user":    userAndExc,
//			"balance": 0,
//		})
//	}
//}
