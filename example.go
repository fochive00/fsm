package fsm


// states
const (
	idle = iota
	walk
	jump
	dead
)

// inputs
const (
	s = iota
	w
	space
)

func main() {
	stateIdle := NewStateBuilder().
		SetStateID(idle).
		

}