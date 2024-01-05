package expirenza

type EventHandler interface {
	Type() Operation
	Handle(*Expirenza, interface{})
}

type EventInterfaceProvider interface {
	Type() Operation
	New() interface{}
}

var registeredInterfaceProviders = map[Operation]EventInterfaceProvider{}

func registerInterfaceProvider(eh EventInterfaceProvider) {
	if _, ok := registeredInterfaceProviders[eh.Type()]; ok {
		return
	}
	registeredInterfaceProviders[eh.Type()] = eh
	return
}

type eventHandlerInstance struct {
	eventHandler EventHandler
}

func (e *Expirenza) addEventHandler(eventHandler EventHandler) func() {
	e.handlersMu.Lock()
	defer e.handlersMu.Unlock()

	if e.handlers == nil {
		e.handlers = map[Operation][]*eventHandlerInstance{}
	}

	ehi := &eventHandlerInstance{eventHandler}
	e.handlers[eventHandler.Type()] = append(e.handlers[eventHandler.Type()], ehi)

	return func() {
		e.removeEventHandlerInstance(eventHandler.Type(), ehi)
	}
}

// addEventHandler adds an event handler that will be fired the next time
// the Discord WSAPI matching eventHandler.Type() fires.
func (e *Expirenza) addEventHandlerOnce(eventHandler EventHandler) func() {
	e.handlersMu.Lock()
	defer e.handlersMu.Unlock()

	if e.onceHandlers == nil {
		e.onceHandlers = map[Operation][]*eventHandlerInstance{}
	}

	ehi := &eventHandlerInstance{eventHandler}
	e.onceHandlers[eventHandler.Type()] = append(e.onceHandlers[eventHandler.Type()], ehi)

	return func() {
		e.removeEventHandlerInstance(eventHandler.Type(), ehi)
	}
}

func (e *Expirenza) AddHandler(handler interface{}) func() {
	eh := e.handlerForInterface(handler)

	if eh == nil {
		return func() {}
	}

	return e.addEventHandler(eh)
}

func (e *Expirenza) AddHandlerOnce(handler interface{}) func() {
	eh := e.handlerForInterface(handler)

	if eh == nil {
		e.log("Invalid handler type, handler will never be called")
		return func() {}
	}

	return e.addEventHandlerOnce(eh)
}

func (e *Expirenza) removeEventHandlerInstance(t Operation, ehi *eventHandlerInstance) {
	e.handlersMu.Lock()
	defer e.handlersMu.Unlock()

	handlers := e.handlers[t]
	for i := range handlers {
		if handlers[i] == ehi {
			e.handlers[t] = append(handlers[:i], handlers[i+1:]...)
		}
	}

	onceHandlers := e.onceHandlers[t]
	for i := range onceHandlers {
		if onceHandlers[i] == ehi {
			e.onceHandlers[t] = append(onceHandlers[:i], onceHandlers[i+1:]...)
		}
	}
}

func (e *Expirenza) handle(t Operation, i interface{}) {
	for _, eh := range e.handlers[t] {
		go eh.eventHandler.Handle(e, i)
	}

	if len(e.onceHandlers[t]) > 0 {
		for _, eh := range e.onceHandlers[t] {
			go eh.eventHandler.Handle(e, i)
		}
		e.onceHandlers[t] = nil
	}
}

func (e *Expirenza) handleEvent(t Operation, i interface{}) {
	e.handlersMu.RLock()
	defer e.handlersMu.RUnlock()

	e.handle(interfaceEventType, i)
	e.handle(t, i)
}
