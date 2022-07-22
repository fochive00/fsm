package fsm

// Builder for FSM
type FSMBuilder struct {
	inner FSM
}

func newFSMBuilder() *FSMBuilder {
	fsm := FSM {
		currentStateID: -1,
		states: make(map[StateID]State),
	}

	return &FSMBuilder {inner: fsm}
}

func (builder *FSMBuilder) SetCurrentState(stateID StateID) *FSMBuilder {
	builder.inner.currentStateID = stateID
	return builder
}

func (builder *FSMBuilder) AddState(state State) *FSMBuilder {
	builder.inner.states[state.stateID] = state
	return builder
}

func (builder *FSMBuilder) Build() FSM {
	return builder.inner
}


// Builder for State
type StateBuilder struct {
	inner State
}

func newStateBuilder() *StateBuilder {
	state := State {
		stateID: -1,
		transition: make(map[Input]Outcome),
	}

	return &StateBuilder {inner: state}
}

func (builder *StateBuilder) SetStateID(stateID StateID) *StateBuilder {
	builder.inner.stateID = stateID
	return builder
}

func (builder *StateBuilder) AddTransition(in Input, out Outcome) *StateBuilder {
	builder.inner.transition[in] = out
	return builder
}

func (builder *StateBuilder) Build() State {
	return builder.inner
}