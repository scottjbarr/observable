package observable

import (
	"sync"
)

// Observable manages registration, deregistration and sending notifications
// to Observers.
type Observable struct {
	observers map[string]Observer
	mu        sync.Mutex
}

// NewObservable returns a new Observable.
func NewObservable() Observable {
	return Observable{
		observers: map[string]Observer{},
		mu:        sync.Mutex{},
	}
}

// Register adds an Observer to this Observable.
func (o Observable) Register(observer Observer) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.observers[observer.Identifier()] = observer
}

// Deregister removes an Observer from this Observable.
func (o Observable) Deregister(observer Observer) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if _, ok := o.observers[observer.Identifier()]; !ok {
		// not found. either not registered, or already deleted.
		return
	}

	delete(o.observers, observer.Identifier())
}

// Notify sends the given message to all registered Observers.
func (o Observable) Notify(message interface{}) {
	o.mu.Lock()
	defer o.mu.Unlock()

	for _, observer := range o.observers {
		observer.Update(message)
	}
}
