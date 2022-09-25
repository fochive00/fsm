package fsm

type EventID string

type EventData interface {
	StateID() StateID
}

type Event interface {
	EventID() EventID
	EventData() EventData
}
