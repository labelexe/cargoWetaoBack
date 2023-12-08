package user

import (
	"Wetao/internal/api_back/app/bootstrap/observer"
	orderModel "Wetao/internal/api_back/app/http/models/order"
	"Wetao/internal/database"
	"crypto/rand"
	"encoding/hex"
	"gorm.io/gorm"
	"log"
)

// AuthData представляет структуру данных для аутентификации Telegram
type AuthData struct {
	AuthDate int64  `json:"auth_date"`
	Hash     string `json:"hash"`
}

// Users представляет структуру данных для пользователя
type Users struct {
	Id int64 `json:"-"`
	// Define your User struct fields here
	TelegramID int64 `json:"telegram_id"`
	//Email      string
	Username   string              `json:"username"`
	IsActive   bool                `json:"is_active"`
	FirstName  string              `json:"first_name"`
	LastName   string              `json:"last_name"`
	Status     string              `json:"status"`
	AuthToken  string              `json:"-" gorm:"auth_token"`
	IsPremium  bool                `json:"is_premium"`
	PhotoURL   string              `json:"photo_url"`
	UserOrders []orderModel.Orders `json:"user_orders" gorm:"foreignKey:UserID;references:TelegramID"`
	//
	UserOrdersInfo struct {
		PendingCount      int64 `json:"pending_count"`
		AcceptedSendCount int64 `json:"sending_accepted"`
	} `json:"user_orders_info" gorm:"-"`

	// Add other fields as needed
}

func generateRandomToken() string {
	tokenBytes := make([]byte, 20)
	rand.Read(tokenBytes)
	return hex.EncodeToString(tokenBytes)
}

func (u *Users) AfterCreate(tx *gorm.DB) error {
	//send new event
	observer.NotifyObserver("app.user.new_customer", u)
	return nil
}

// GetUserOrdersInfo метод для получения информации о заказах пользователя
func (u *Users) GetUserOrdersInfo(db *gorm.DB) error {

	var pendingCount, sendingAcceptedCount int64
	db.Model(&u).
		Joins("JOIN orders ON users.telegram_id = orders.user_id").
		Where("orders.user_id = ? AND orders.status = 'pending_accepted'", u.TelegramID).
		Count(&pendingCount)

	db.Model(&u).
		Joins("JOIN orders ON users.telegram_id = orders.user_id").
		Where("orders.user_id = ? AND orders.status = 'sending_accepted'", u.TelegramID).
		Count(&sendingAcceptedCount)

	// Запись значений в структуру user.UserOrdersInfo
	u.UserOrdersInfo.PendingCount = pendingCount
	u.UserOrdersInfo.AcceptedSendCount = sendingAcceptedCount

	return nil
}

func SaveUserData(user Users) error {
	db := database.GetDB()
	u := Users{TelegramID: user.TelegramID, Username: user.Username}
	result := db.Create(&u)
	return result.Error
}

func GetUserData(TelegramID int64) (Users, error) {
	db := database.GetDB()
	var user Users
	result := db.First(&user, TelegramID)
	return user, result.Error
}

func GetUserByTelegramID(telegramID int) *Users {
	// Implement the logic for fetching user by Telegram ID
	// ...
	db := database.GetDB()

	// Return dummy user data for testing
	var user Users
	if err := db.First(&user, "telegram_id = ?", telegramID).Error; err != nil {
		return nil
	}
	return &user
}

func CreateUser(telegramID int64, username string, firstName string, lastName string) *Users {
	// creating a new user
	db := database.GetDB()

	authToken := generateRandomToken()

	// Return dummy user data for testing
	newUser := &Users{
		TelegramID: telegramID,
		Username:   username,
		IsActive:   true,
		FirstName:  firstName,
		LastName:   lastName,
		//Password:   authToken, // Replace with proper password hashing
		Status:    "Active",
		AuthToken: authToken,
		IsPremium: false,
		// Populate other fields as needed
	}

	if err := db.Create(newUser).Error; err != nil {
		log.Println("Error creating user:", err)
		return nil
	}

	return newUser
}
