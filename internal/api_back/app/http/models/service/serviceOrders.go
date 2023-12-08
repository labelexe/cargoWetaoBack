package service

//
//import (
//	"database/sql"
//	"gorm.io/gorm"
//	"retryh-billing-go/app/database"
//	"retryh-billing-go/app/models/service"
//	"time"
//)
//
//// UserOrderServices Заказы
//type UserOrderServices struct {
//	ID          uint              `json:"-" gorm:"primaryKey"`
//	UserID      uint              `json:"-" gorm:"user_id"`
//	TotalAmount float64           `json:"total_amount" gorm:"total_amount"`
//	OrderParams string            `json:"-" gorm:"order_params"`
//	PromoCode   string            `json:"-" gorm:"promo_code"`
//	Status      string            `json:"status" gorm:"status"`
//	CreatedAt   time.Time         `json:"-"`
//	UpdatedAt   time.Time         `json:"-"`
//	Services    []OrderedServices `json:"-" gorm:"foreignKey:OrderID"`
//}
//
//// OrderedServices Заказанные услуги
//type OrderedServices struct {
//	ID                  uint                    `json:"-" gorm:"primaryKey"`
//	TariffID            uint                    `json:"-" gorm:"tariff_id"`
//	ServiceID           uint                    `json:"-" gorm:"service_id"`
//	ServiceInstructions string                  `json:"service_about"`
//	ServicePrice        float64                 `json:"service_price"`
//	OrderID             uint                    `json:"-" gorm:"order_id"`
//	OrderStatus         string                  `json:"service_status"`
//	Order               UserOrderServices       `json:"order_info" gorm:"foreignKey:OrderID"`
//	Service             service.Services        `json:"service_info" gorm:"foreignKey:ServiceID"`
//	TariffsServices     service.TariffsServices `json:"service_tariff" gorm:"foreignKey:TariffID"`
//	// Другие поля услуги
//}
//
////// OrderedServices Заказанные услуги
////type OrderedServices struct {
////	Id        uint `json:"-" gorm:"id, primaryKey"`
////	ServiceID uint `json:"-" gorm:"service_id"` // ID заказанной услуги
////	//TariffID            uint    `json:"-" gorm:"tariff_id"`
////	ServiceInstructions string  `json:"service_about" gorm:"order_id"` // Нужна админам для инструкции на заказанную услугу
////	ServicePrice        float64 `json:"-" gorm:"service_price"`        //Текущая цена данной заказанной услуги
////	//
////	OrderID     uint   `json:"-" gorm:"order_id"`              // ID заказа
////	OrderStatus string `json:"service_status" gorm:"order_id"` // Статус заявки на данную услугу
////	//
////	Orders []UserOrderServices `json:"order_info" gorm:"foreignKey:OrderID"`
////	// О текущем сервисе
////	Service service.Services `json:"service_info" gorm:"foreignKey:ServiceID"`
////	//
////	//ServiceTariff service.TariffsServices `json:"tariff_info" gorm:"foreignKey:TariffID"`
////}
////
////// UserOrderServices Заказы
////type UserOrderServices struct {
////	Id          uint    `json:"-" gorm:"id, primaryKey"`
////	UserID      uint    `json:"-" gorm:"user_id"`
////	TotalAmount float64 `json:"total_amount" gorm:"total_amount"`
////	PromoCode   string  `json:"-" gorm:"promo_code"`
////	//OrderParams map[string]interface{} `json:"-" gorm:"order_params, type:jsonb"`
////	Status    string    `json:"status" gorm:"status"`
////	CreatedAt time.Time `json:"-"`
////	UpdatedAt time.Time `json:"-"`
////	//
////	Services []OrderedServices `json:"services" gorm:"foreignKey:OrderID"`
////
////	// Другие поля заявки
////}
//
//// createNewServiceOrder - Создание нового пользовательского заказа по тарифу на
//func createNewServiceOrder(userId int64, tariffId uint) error {
//	//TODO: За тем найти желаемую услугу, и тариф если они есть в наличие
//	//TODO: Мы должны свериться по сделанным параметрам, и рассчитать конечную сумму заказа!
//	//TODO: Далее нужно чтобы мы проверили баланс пользователя
//
//	//availBalanceByParams = 00.00
//
//	// Расчет суммы заказанных услуг
//
//	//paramsService
//
//	return nil
//
//}
//
//// GetServiceOrdersByUserId - Получение заказов пользователя и его услуги
//func GetServiceOrdersByUserId(userId uint) []UserOrderServices {
//	if userId == 0 {
//		return []UserOrderServices{}
//	}
//
//	//User-id not nil
//	var serviceList []UserOrderServices
//	// Fetch services
//	if database.
//		GetDB().
//		Select("*").
//		Where("user_id = @user_id", sql.Named("user_id", userId)).
//		Preload("Services").
//		Preload("Services.Order").
//		Preload("Services.Service").
//		Find(&serviceList).Error != nil {
//		return serviceList
//	} else {
//		return serviceList
//	}
//}
//
//// GetOrdersByUserId - Получение заказов пользователя и его услуги
//func GetOrdersByUserId(userId uint) []OrderedServices {
//	if userId == 0 {
//		return []OrderedServices{}
//	}
//
//	//User-id not nil
//	var serviceList []OrderedServices
//	var userOrderServices []UserOrderServices
//
//	// Найти заказ пользователя по полю "user_id"
//	database.
//		GetDB().
//		Select("*").
//		Preload("Services", func(db *gorm.DB) *gorm.DB {
//			return db.
//				Preload("Order").
//				Preload("Service").
//				Preload("TariffsServices")
//		}).
//		Where("user_id = @user_id", sql.Named("user_id", userId)).
//		//Preload("Services.TariffsServices").
//		//Omit("Services.Order.Services").
//		Find(&userOrderServices)
//	//
//	for _, userOrder := range userOrderServices {
//		for _, orderService := range userOrder.Services {
//			//fmt.Println(orderService)
//			//newOrderService := UserOrderServices{}
//			//"Order.Services"
//			serviceList = append(serviceList, orderService)
//		}
//	}
//
//	return serviceList
//}
