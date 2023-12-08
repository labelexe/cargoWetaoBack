package event

import (
	"github.com/rs/zerolog/log"
)

// NewCustomerListener представляет конкретный EventListener
type NewCustomerListener struct {
	Name string
}

func NewCustomerEvent() NewCustomerListener {
	return NewCustomerListener{}
}

// HandleEvent - метод для обработки событий
func (e *NewCustomerListener) HandleEvent(data interface{}) {
	//fmt.Printf("EventListener %s обработал событие: %v\n", e.Name, data)
	log.Print("Ново-рег - ", data)
}
