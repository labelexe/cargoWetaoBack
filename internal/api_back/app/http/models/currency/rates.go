package currency

import (
	"Wetao/internal/database"
)

func (CurRate) TableName() string {
	return "currency_rates"
}

// CurRate Model Table
type CurRate struct {
	ID           uint    `json:"-" gorm:"primaryKey"`
	Currency     string  `json:"-" gorm:"not null"`
	DirCurrency  string  `json:"currency" gorm:"not null"`
	DirCurSymbol string  `json:"currency_sym" gorm:"not null"`
	Value        float64 `gorm:"not null"`
}

// GetManyCurrencyDirRates Получение сразу по нескольким направлениями которые мы собрали
func GetManyCurrencyDirRates(manyDirCurrency []string) ([]CurRate, []string, error) {
	var currencyRates []CurRate

	firstData := database.GetDB().
		Select("*").
		//Where(&fields).
		Where("currency = ? AND dir_currency IN ?", "RUB", manyDirCurrency).
		Find(&currencyRates)
	//Scan(&currencyRates)

	return currencyRates, manyDirCurrency, firstData.Error
}

// GetExcRateByDirCurrency - Получение по одному направлению
func GetExcRateByDirCurrency(dirCurrency string) (CurRate, string, error) {
	var currencyRates CurRate
	firstData := database.GetDB().
		Select("*").
		Where("currency = ? AND dir_currency = ?", "RUB", dirCurrency).
		First(&currencyRates)

	return currencyRates, dirCurrency, firstData.Error
}
