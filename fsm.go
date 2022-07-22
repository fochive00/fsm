package fsm

type Input int
type StateID int
type Action func()

type Outcome struct {
	nextStateID StateID
	action		Action
}

type State struct {
	stateID StateID
	transition map[Input]Outcome
}

type FSM struct {
	states			map[StateID]State
	currentStateID 	StateID
	
}

func (fsm *FSM)TakeInput(in Input) {
	currentState := fsm.states[fsm.currentStateID]
	Outcome := currentState.transition[in]

	// TODO
	// Action

	fsm.currentStateID = Outcome.nextStateID
}

