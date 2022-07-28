package fsm

import (
	"sync"
)

type FSMPool struct {
	pool            *sync.Pool
	transitionTable TransitionTable
}

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

func (fsmPool *FSMPool) Put(fsm *FSM) {
	fsmPool.pool.Put(fsm)
}

func (fsmPool *FSMPool) Get() *FSM {
	// Get from pool and do type assertion
	fsm := fsmPool.pool.Get().(*FSM)

	// Reset initialized value
	fsm.initialized = false
	return fsm
}
