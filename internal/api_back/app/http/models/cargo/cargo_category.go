package cargo

import (
	"Wetao/internal/database"
)

func (Category) TableName() string {
	return "cargo_category"
}

// Category Категории (карго)
type Category struct {
	ID            uint   `json:"-" gorm:"primaryKey"`
	CargoType     uint   `json:"-" gorm:"cargo_type"`
	CargoWithType Type   `json:"type" gorm:"foreignKey:id"`
	Name          string `json:"name" gorm:"name"`
	Slug          string `json:"slug"`
	//CreatedAt   time.Time         `json:"-"`
	//UpdatedAt   time.Time         `json:"-"`
	//Services    []OrderedServices `json:"-" gorm:"foreignKey:OrderID"`
}

// GetAllCargoCategoriesAndTypes Получение категорий и вложенные типы (Карго)
func GetAllCargoCategoriesAndTypes() ([]Category, error) {
	var cargoCategories []Category

	dbData := database.GetDB().
		Select("*").
		Preload("CargoWithType").
		//Where(&fields).
		//Where("currency = ? AND dir_currency IN ?", "RUB", manyDirCurrency).
		Find(&cargoCategories)

	return cargoCategories, dbData.Error
}

func GetBySlug() {
	
}
