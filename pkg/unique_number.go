package pkg

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

var (
	uniqueNumbers      = make(map[int64]bool)
	uniqueNumbersMutex sync.Mutex
)

// generateUniqueNumber генерирует уникальное число в диапазоне 1904 – 1 G
func generateUniqueNumber() int {
	uniqueNumbersMutex.Lock()
	defer uniqueNumbersMutex.Unlock()

	var maxNumber = new(big.Int)
	maxNumber, _ = maxNumber.SetString("99999", 10) // Максимальное число в диапазоне 1904 – 1 G

	var number *big.Int
	for {
		bytes, err := rand.Int(rand.Reader, maxNumber)
		if err != nil {
			panic(err)
		}

		number = new(big.Int).Add(bytes, big.NewInt(1904))
		if !uniqueNumbers[number.Int64()] {
			uniqueNumbers[number.Int64()] = true
			break
		}
	}

	return int(number.Int64())
}

// GenerateUniqueID генерирует уникальный идентификатор в формате "XXXX – Y G"
func GenerateUniqueID() string {
	uniqueNumber := generateUniqueNumber()
	uniqueID := fmt.Sprintf("%d – %d G", uniqueNumber, 1)
	return uniqueID
}
