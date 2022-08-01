package fsm

import (
	"sync"
)

type FSMPool struct {
	pool            *sync.Pool
	transitionTable TransitionTable
}

// This is a pool wrapper.
// We need to initialize FSM pool with a transition table.
// Every FSM from this pool will share the same transition table.
func NewFSMPool(transitionTable TransitionTable) *FSMPool {
	fsmPool := &FSMPool{
		pool: &sync.Pool{
			// The Pool's New function should generally only return pointer
			// types, since a pointer can be put into the return interface
			// value without an allocation:
			New: func() any {
				return &FSM{
					transitionTable: transitionTable,
				}
			},
		},
		transitionTable: transitionTable,
	}

	return fsmPool
}

// Put the FSM back to pool for next time we use it.
func (fsmPool *FSMPool) Put(fsm *FSM) {
	fsmPool.pool.Put(fsm)
}

// Get a FSM from pool and mark it as uninitialized.
func (fsmPool *FSMPool) Get() *FSM {
	// Get from pool and do type assertion
	fsm := fsmPool.pool.Get().(*FSM)

	// Reset initialized value
	fsm.initialized = false
	return fsm
}
