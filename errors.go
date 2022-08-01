package fsm

import "fmt"

type InvalidInputError struct {
	state State
	input Input
}

func (err InvalidInputError) Error() string {
	return fmt.Sprintf("input invalid in current state.  (State: %v, Input: %v)", err.state, err.input)
}

// type StateNotInitializedError struct{}

// func (err StateNotInitializedError) Error() string {
// 	return "FSM not initialized."
// }

// type ImpossibleStateError State

// func (err ImpossibleStateError) Error() string {
// 	return fmt.Sprintf("FSM in impossible state.  (State: %d)", err.String())
// }
