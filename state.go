// Design Pattern: FlyWeight
//
// Members:
// FlyWeight 					- State
// ConcreteFlyWeight			- xxxState
// UnSharedConcreteFlyWeight	-
// FlyWeightFactory 			- StateFactory
// FlyWeightFactoryBuilder		- StateFactoryBuilder
// Client						-

package fsm

type StateID string

// FlyWeight
type State interface {
	Handle(event Event)
}

// FlyWeight Factory
type StateFactory interface {
	GetState(stateID StateID) (State, bool)
}

type stateFactory struct {
	stateMap map[StateID]State
}

func (f *stateFactory) GetState(stateID StateID) (State, bool) {
	state, ok := f.stateMap[stateID]

	return state, ok
}

// Flyweight Factory Builder
type StateFactoryBuilder interface {
	AddState(stateID StateID, state State) StateFactoryBuilder
	Build() StateFactory
}

type stateFactoryBuilder struct {
	stateMap map[StateID]State
}

func NewStateFactoryBuilder() StateFactoryBuilder {
	return &stateFactoryBuilder{
		stateMap: make(map[StateID]State),
	}
}

func (fb *stateFactoryBuilder) AddState(stateID StateID, state State) StateFactoryBuilder {
	_, ok := fb.stateMap[stateID]
	if ok {
		panic("state already exists")
	}

	fb.stateMap[stateID] = state

	return fb
}

func (fb *stateFactoryBuilder) Build() StateFactory {
	return &stateFactory{
		stateMap: fb.stateMap,
	}
}
