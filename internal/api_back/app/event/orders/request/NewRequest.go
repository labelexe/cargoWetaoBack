package request

import (
	orderModel "Wetao/internal/api_back/app/http/models/order"
	"Wetao/internal/tg_back/order/request"
	"github.com/rs/zerolog/log"
)

// NewRequestOrderListener представляет конкретный EventListener
type NewRequestOrderListener struct {
	Name string
}

func RegUserEvent() NewRequestOrderListener {
	return NewRequestOrderListener{}
}

// HandleEvent - метод для обработки событий
func (e *NewRequestOrderListener) HandleEvent(data interface{}) {
	//fmt.Printf("EventListener %s обработал событие: %v\n", e.Name, data)
	log.Print("Поступил новый заказ! -", data)
	request.TelegramSendMsgNewOrderReq(data.(orderModel.Orders))
}
