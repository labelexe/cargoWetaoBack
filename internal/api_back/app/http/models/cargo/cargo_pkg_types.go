package cargo

import "Wetao/internal/database"

func (PackageTypes) TableName() string {
	return "cargo_pkg_types"
}

// PackageTypes Категории (карго)
type PackageTypes struct {
	ID   uint   `json:"-" gorm:"primaryKey"`
	Name string `json:"name" gorm:"name"`
	Slug string `json:"slug"`
	//CreatedAt   time.Time         `json:"-"`
	//UpdatedAt   time.Time         `json:"-"`
	//Services    []OrderedServices `json:"-" gorm:"foreignKey:OrderID"`
}

// GetAllCargoPkgTypes Получение типа упаковки (Карго)
func GetAllCargoPkgTypes() ([]PackageTypes, error) {
	var cargoCategories []PackageTypes

	dbData := database.GetDB().
		Select("*").
		//Preload("CargoWithType").
		//Where(&fields).
		//Where("currency = ? AND dir_currency IN ?", "RUB", manyDirCurrency).
		Find(&cargoCategories)

	return cargoCategories, dbData.Error
}
