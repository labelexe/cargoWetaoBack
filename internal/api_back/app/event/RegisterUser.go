package event

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// RegUserListener представляет конкретный EventListener
type RegUserListener struct {
	Name string
}

func RegUserEvent() RegUserListener {
	return RegUserListener{}
}

// HandleEvent - метод для обработки событий
func (e *RegUserListener) HandleEvent(data interface{}) {
	//fmt.Printf("EventListener %s обработал событие: %v\n", e.Name, data)
	log.Print("Ново-рег - ", data.(fiber.Map)["пример данных"])
}
