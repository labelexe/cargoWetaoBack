package order

import (
	"Wetao/internal/api_back/app/http/models/cargo"
	"Wetao/internal/database"
	"database/sql"
	"fmt"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type WeightSizes struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//func (Orders) TableName() string {
//	return "orders"
//}

type Orders struct {
	ID     uint   `json:"-" gorm:"primarykey"`
	Idx    string `json:"idx" gorm:"idx"`
	UserID int64  `json:"-" gorm:"user_id"`
	//User             userModel.Users `json:"user" gorm:"foreignKey:TelegramID"`
	TypeBuyback      string             `json:"type_buyback"`
	Slug             string             `json:"slug"`
	CatID            uint               `json:"-" gorm:"cat_id"`
	PkgID            uint               `json:"-" gorm:"pkg_id"`
	OrderCategory    cargo.Category     `json:"cat" gorm:"foreignKey:id;references:cat_id"`
	OrderPackage     cargo.PackageTypes `json:"pkg" gorm:"foreignKey:id;references:pkg_id"`
	WeightSizes      WeightSizes        `json:"weight_sizes" gorm:"serializer:json"`
	TotalAmountPrice float64            `json:"total_amount_price"`
	ProductLinkURL   string             `json:"product_link_url"`
	Status           string             `json:"status"`
	CreatedAt        time.Time          `json:"created_at"`          // Exported field
	UpdatedAt        time.Time          `json:"-" gorm:"updated_at"` // Exported field
	DeletedAt        *gorm.DeletedAt    `gorm:"index" json:"-"`
}

// GenerateUniqueSlug generates a unique slug starting with "webuy_"
func GenerateUniqueSlug(prefix string) string {
	uuidStr := uuid.NewV4()
	randomString := strings.Replace(uuidStr.String(), "-", "", -1)
	return prefix + "_" + randomString
}

// InsertNewRequestByOrder - Создание новой заявки
func InsertNewRequestByOrder(newOrder *Orders) (Orders, error) {
	db := database.GetDB()
	var resultNewOrder Orders

	// Create the new order in the database
	if err := db.Create(&newOrder).Error; err != nil {
		fmt.Println("Error creating the order:", err)
		return Orders{}, err
	}

	// Retrieve the created order from the database
	if err := db.First(&resultNewOrder, newOrder.ID).Error; err != nil {
		fmt.Println("Error retrieving the order:", err)
		return Orders{}, err
	}

	return resultNewOrder, nil
}

func FilterOrderByUserId(userTgId int64) ([]Orders, error) {
	db := database.GetDB()
	var resultOrders []Orders

	// Retrieve the created order from the database
	if err := db.
		Select("*").
		Where("user_id = @user_id",
			sql.Named("user_id", userTgId),
		).
		Find(&resultOrders).Error; err != nil {
		fmt.Println("Error retrieving the order:", err)
		return []Orders{}, err
	}

	return resultOrders, nil
}
