package state

import (
	"github.com/fochive00/examples/event"
	"github.com/fochive00/fsm"
	"github.com/fochive00/fsm/examples/event"
)

const IdleStateID = "idle"

type IdleState struct {
	stateID        fsm.StateID
	handlerFactory fsm.HandlerFactory
}

func NewIdleState() IdleState {
	state := IdleState{
		stateID: IdleStateID,
	}

	handlerFactory := fsm.NewHandlerFactoryBuilder().
		AddHandler(event.AttackEventId, state.handleAttack).
		AddHandler(event.MoveEventID, state.handleMove).
		AddHandler(event.StopEventID, nil).
		AddHandler(event.takeDamage, nil).
		Build()

	state.handlerFactory = handlerFactory

	return state
}

func (state *IdleState) StateID() fsm.StateID {
	return state.stateID
}

func (state *IdleState) HandlerFactory() fsm.HandlerFactory {
	return state.handlerFactory
}

func (state *IdleState) handleAttack(event fsm.Event) {

}

func (state *IdleState) handleMove(event fsm.Event) {

}
