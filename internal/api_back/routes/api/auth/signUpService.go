package auth

//
//import (
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	userModel "retryh-billing-go/app/models/user"
//	"retryh-billing-go/pkg/events"
//)
//
//type UserSignUp struct {
//	User     string `json:"user"`
//	Password string `json:"password"`
//}
//
//// TODO: Register!
//
//// SignUp @Router /api/v1/auth/signup
//func SignUp(c *fiber.Ctx) error {
//	user := new(UserSignUp)
//	if err := c.BodyParser(user); err != nil {
//		fmt.Println("error = ", err)
//		return c.SendStatus(200)
//	}
//
//	//Тело-запроса
//	userReqs := new(UserSignIn)
//	if err := c.BodyParser(userReqs); err != nil {
//		fmt.Println("error = ", err)
//		return c.SendStatus(200)
//	}
//
//	createUser, err := userModel.CreateUser(userReqs.User, userReqs.Password)
//	if err != nil {
//		return err
//	}
//
//	fmt.Println(createUser)
//
//	//var users []UserSignUp
//	//
//	//for i := 0; i < 900000; i++ {
//	//	users = append(users, UserSignUp{
//	//		User:     fmt.Sprintf("user%dname", i),
//	//		Password: fmt.Sprintf("2432%drwerew", i),
//	//	})
//	//}
//
//	events.NewPushEventToBroadcast("newRegisterUser", createUser)
//
//	return c.JSON(fiber.Map{"success": true})
//}
//
//func SignUpView(c *fiber.Ctx) error {
//	return c.Render("auth/signup", fiber.Map{
//		"Title": "Регистрация",
//	}, "auth/layout/auth")
//}
