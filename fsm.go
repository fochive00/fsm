package fsm

import "sync"

type Input int
type State int
type Action func() Input

type Output struct {
	nextState State
	action    Action
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

func (fsm *FSM) TakeInput(input Input) error {
	fsm.Lock()
	defer fsm.Unlock()

	if !fsm.initialized {
		return StateNotInitializedError{}
	}

	transition, ok := fsm.transitionTable[fsm.currentState]

	if !ok {
		return ImpossibleStateError(fsm.currentState)
	}

	output, ok := transition[input]

	if !ok {
		return InvalidInputError{fsm.currentState, input}
	}

	// TODO chain Action
	output.action()

	fsm.currentState = output.nextState

	return nil
}

func (fsm *FSM) InitState(state State) {
	fsm.Lock()
	defer fsm.Unlock()

	fsm.currentState = state
	fsm.initialized = true
}
