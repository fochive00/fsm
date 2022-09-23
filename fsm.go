package fsm

import (
	"errors"
	"log"
	"sync"
)

// type Input int
// type State int
type Input interface {
	String() string
}

type Action func() Input

type Output struct {
	NextState State
	Action    Action
}

// A transition table indicates that what next state will be
// and action will take for each input of each current state.
type TransitionTable map[State]map[Input]Output

type FSM struct {
	sync.Mutex
	initialized     bool
	currentState    State
	transitionTable TransitionTable
}

// Transitions from current state to the next state
// and take necessary actions.
func (fsm *FSM) Spin(input Input) error {
	// fsm.Lock()
	// defer fsm.Unlock()

	// Check if the FSM current state is initialized
	if !fsm.initialized {
		return errors.New("FSM not initialized")
	}

	transition, ok := fsm.transitionTable[fsm.currentState]

	if !ok {
		// Error: FSM in impossible state.
		log.Panicf("FSM in impossible state.  (State: %s)", fsm.currentState.String())
	}

	output, ok := transition[input]

	if !ok {
		return InvalidInputError{fsm.currentState, input}
	}

	// do nothing when action is nil
	if output.Action != nil {
		output.Action()
	}

	fsm.currentState = output.NextState

	return nil
}

// Every time you get a new FSM from pool,
// you need to initialize the current state
func (fsm *FSM) InitState(state State) {
	// fsm.Lock()
	// defer fsm.Unlock()

	fsm.currentState = state
	fsm.initialized = true
}

func (fsm *FSM) GetState() State {
	return fsm.currentState
}
