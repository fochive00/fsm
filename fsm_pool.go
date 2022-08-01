package fsm

import (
	"sync"
)

type FSMPool struct {
	pool            *sync.Pool
	transitionTable TransitionTable
}

// simple counter for benchmark
// type atomInt struct {
// 	sync.Mutex
// 	value int
// }
// var counter = aint{value: 0}

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
				// counter setup
				// counter.Lock()
				// defer counter.Unlock()
				// fmt.Printf("New FSM from pool. %d", counter.value)
				// counter.value++

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
	// The User must initialize current state of the FSM that are getting from the pool.
	fsm.initialized = false

	// fmt.Println("Get a FSM from pool.")
	return fsm
}
