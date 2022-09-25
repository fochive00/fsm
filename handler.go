package fsm

type Handler func(event Event)

// Simple Factory
type HandlerFactory interface {
	GetHandler(eventID EventID) (Handler, bool)
}

type handlerFactory struct {
	handlerMap map[EventID]Handler
}

func (f *handlerFactory) GetHandler(eventID EventID) (Handler, bool) {
	handler, ok := f.handlerMap[eventID]
	return handler, ok
}

// Factory Builder
type HandlerFactoryBuilder interface {
	AddHandler(eventID EventID, handler Handler) HandlerFactoryBuilder
	Build() HandlerFactory
}

type handlerFactoryBuilder struct {
	handlerMap map[EventID]Handler
}

func NewHandlerFactoryBuilder() HandlerFactoryBuilder {
	return &handlerFactoryBuilder{
		handlerMap: make(map[EventID]Handler),
	}
}

func (fb *handlerFactoryBuilder) AddHandler(eventID EventID, handler Handler) HandlerFactoryBuilder {
	_, ok := fb.handlerMap[eventID]
	if ok {
		panic("handler already exists")
	}

	fb.handlerMap[eventID] = handler

	return fb
}

func (fb *handlerFactoryBuilder) Build() HandlerFactory {
	return &handlerFactory{
		handlerMap: fb.handlerMap,
	}
}
