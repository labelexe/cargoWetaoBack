package cargo

import (
	"Wetao/internal/database"
)

func (WghSizes) TableName() string {
	return "cargo_wgh_sizes"
}

// WghSizes Размеры грузов (карго)
type WghSizes struct {
	ID   uint   `json:"-" gorm:"primaryKey"`
	Name string `json:"name" gorm:"name"`
	//Slug             string `json:"slug"`
	DensityRangeFrom int     `json:"-"`
	DensityRangeTo   int     `json:"-"`
	Price            float64 `json:"price" gorm:"price"`
	CargoTypeID      int     `json:"-" gorm:"cargo_type_id"`
	CargoCatID       int     `json:"-" gorm:"cargo_cat_id"`
}

// GetAllWghSizesDistName Получение всех размеров груза (Карго) (Группировка)
func GetAllWghSizesDistName() ([]WghSizes, error) {
	var cargoWghSizes []WghSizes

	dbData := database.GetDB().
		//Model(&WghSizes{}).
		Select("DISTINCT ON (name) *").
		Order("name").
		Find(&cargoWghSizes)

	//fmt.Println(cargoWghSizes)

	return cargoWghSizes, dbData.Error
}
