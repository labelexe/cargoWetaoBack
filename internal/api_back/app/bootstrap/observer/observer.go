package observer

import (
	"fmt"
	"sync"
)

// EventListener interface defines the method to handle events.
type EventListener interface {
	HandleEvent(data interface{})
}

// Observer interface defines the update method that observers must implement.
type Observer interface {
	EventListener
}

// ObserverRegistry maintains a global registry of observers.
type ObserverRegistry struct {
	observers map[string]Observer
	mu        sync.Mutex
}

var globalObserverRegistry ObserverRegistry

// RegisterObserver adds a new observer to the global registry.
func RegisterObserver(name string, observer Observer) {
	globalObserverRegistry.mu.Lock()
	defer globalObserverRegistry.mu.Unlock()

	if globalObserverRegistry.observers == nil {
		globalObserverRegistry.observers = make(map[string]Observer)
	}

	globalObserverRegistry.observers[name] = observer
}

// UnregisterObserver removes an observer from the global registry.
func UnregisterObserver(name string) {
	globalObserverRegistry.mu.Lock()
	defer globalObserverRegistry.mu.Unlock()

	delete(globalObserverRegistry.observers, name)
}

// NotifyObservers sends an update to the specific observer.
func NotifyObserver(name string, data interface{}) {
	globalObserverRegistry.mu.Lock()
	defer globalObserverRegistry.mu.Unlock()

	if observer, ok := globalObserverRegistry.observers[name]; ok {
		observer.HandleEvent(data)
	} else {
		fmt.Printf("Observer with name %s not found\n", name)
	}
}

// ConcreteObserver is a struct that implements the Observer interface.
type ConcreteObserver struct {
	Name string
}

// HandleEvent prints the updated data.
func (o *ConcreteObserver) HandleEvent(data interface{}) {
	fmt.Printf("Observer %d received an update: %v\n", o.Name, data)
}
