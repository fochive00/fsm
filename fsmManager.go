package fsm

import (
	"sync"

	"github.com/fochive00/fsm/utils"
)

type FsmManager interface {
}

// `Count` indicates that how many fsm are in use.
type fsmManager struct {
	pool            *sync.Pool
	transitionTable TransitionTable
	counter         utils.Counter
}

// This is a pool wrapper.
// We need to initialize FSM pool with a transition table.
// Every FSM from this pool will share the same transition table.
func NewFsmManager(transitionTable TransitionTable) FsmManager {
	manager := &fsmManager{
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
		counter:         utils.NewCounter(),
	}

	return manager
}

// Put the FSM back to pool for next time we use it.
func (manager *fsmManager) Put(fsm *FSM) {
	manager.pool.Put(fsm)

	manager.counter.Decrease()
}

// Get a FSM from pool and mark it as uninitialized.
func (manager *fsmManager) Get() *FSM {
	// Get from pool and do type assertion
	fsm := manager.pool.Get().(*FSM)

	// Reset initialized value
	// The User must initialize current state of the FSM that are getting from the pool.
	fsm.initialized = false

	manager.counter.Increase()

	// fmt.Println("Get a FSM from pool.")
	return fsm
}

func (manager *fsmManager) Count() int64 {
	return manager.counter.Get()
}
