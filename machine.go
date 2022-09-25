package fsm

type Machine interface {
	Handle(event Event)
	StateFactory() StateFactory
}

type machine struct {
	stateFactory StateFactory
}

func (m *machine) Handle(event Event) {
	stateID := event.EventData().StateID()

	state, ok := m.stateFactory.GetState(stateID)
	if !ok {
		panic("state not found")
	}

	eventID := event.EventID()

	handler, ok := state.HandlerFactory().GetHandler(eventID)
	if !ok {
		panic("handler not found")
	}

	if handler != nil {
		handler(event)
	}
}

func (m *machine) StateFactory() StateFactory {
	return m.stateFactory
}

type MachineBuilder interface {
	SetStateFactory(stateFactory StateFactory) MachineBuilder
	Build() Machine
}

type machineBuilder struct {
	stateFactory StateFactory
}

func NewMachineBuilder() *machineBuilder {
	return new(machineBuilder)
}

func (builder *machineBuilder) SetStateFactory(stateFactory StateFactory) MachineBuilder {
	builder.stateFactory = stateFactory

	return builder
}

func (builder *machineBuilder) Build() Machine {
	return &machine{
		stateFactory: builder.stateFactory,
	}
}
