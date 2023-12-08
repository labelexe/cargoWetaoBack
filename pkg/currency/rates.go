package currency

import (
	currencyModel "Wetao/internal/api_back/app/http/models/currency"
	"Wetao/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CurrencyData struct {
	Date         string
	PreviousDate string
	PreviousURL  string
	Timestamp    string
	Valute       map[string]struct {
		ID       string
		NumCode  string
		CharCode string
		Nominal  int
		Name     string
		Value    float64
		Previous float64
	}
}

func parseDailyByCbr() (CurrencyData, error) {
	// Fetch JSON data from the URL
	url := "https://www.cbr-xml-daily.ru/daily_json.js"
	resp, err := http.Get(url)
	if err != nil {
		return CurrencyData{}, err
	}
	defer resp.Body.Close()

	//
	var data CurrencyData

	//
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(fmt.Errorf("currency-data-json err: %d", err))
		return CurrencyData{}, err
	} else {
		return data, nil
	}

}

func updateRates() {
	timestamp := getConsoleTimestamp()

	cbrDaily, err := parseDailyByCbr()
	if err != nil {
		fmt.Println(fmt.Sprintf("R-[%s] Currency-data-http err: %d", timestamp, err))
	}

	// Define the specific currencies to select in a map
	specificCurrencies := map[string]bool{
		"EUR": true,
		"USD": true,
		"AED": true,
		"CNY": true,
	}

	// Iterate through Valute and save currency rates to the database
	for _, currency := range cbrDaily.Valute {
		if specificCurrencies[currency.CharCode] {
			rate := currencyModel.CurRate{
				Currency:    "RUB",
				DirCurrency: currency.CharCode,
				Value:       currency.Value,
			}

			timestamp := getConsoleTimestamp()
			//
			result := database.GetDB().
				Where(&currencyModel.CurRate{DirCurrency: rate.DirCurrency}).
				Assign(rate).
				FirstOrCreate(&rate)
			//
			if result.Error != nil {

				fmt.Println(fmt.Sprintf("R-[%s] Update->currency-rates->db err: %d", timestamp, err))
				//panic(result.Error)
			} else {
				//rate.Value = rate.Value
				result.Save(rate)
				fmt.Println(fmt.Sprintf("R-[%s] Rates updates: to %s -> %s / %f", timestamp, rate.DirCurrency, rate.Currency, rate.Value))
			}
		} else {
		}
	}
}

func UpdateRtByCurAndGet(dirsCurrency []string) ([]currencyModel.CurRate, error) {
	// Обновляем последние курсы и записываем в бд
	updateRates()

	//Запрашиваем из бд дабы удостовериться
	//manyCurrencyDirRates, _, err := currency.GetManyCurrencyDirRates([]string{"USD", "CNY"})
	//if err != nil {
	//	return []CurrencyRate{}, err
	//}
	//return manyCurrencyDirRates, nil
	return nil, nil
}

func getConsoleTimestamp() string {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return formatted
}

//func RunPeriodGetCurrencyRate() {
//	for {
//		// Создаем канал (channel) для отправки сигнала таймеру
//		timer := time.NewTimer(45 * time.Second)
//		updateRates()
//		fmt.Println("----")
//
//		// Block until the timer channel receives a value
//		<-timer.C
//	}
//}
