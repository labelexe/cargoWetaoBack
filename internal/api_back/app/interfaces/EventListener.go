package interfaces

// EventListener представляет интерфейс для обработки событий
type EventListener interface {
	HandleEvent(data interface{})
}
