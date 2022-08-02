package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/fochive00/fsm"
)

type State int
type Input int

func (s State) String() string {
	return fmt.Sprintf("%d", s)
}

func (i Input) String() string {
	return fmt.Sprintf("%d", i)
}

const (
	idle State = iota
	moving
	attacking
	dead
)

const (
	stop Input = iota
	move
	attack
	takeDamage
)

var transitionTable = fsm.TransitionTable{
	idle: map[fsm.Input]fsm.Output{
		stop:       {NextState: idle, Action: nil},
		move:       {NextState: moving, Action: nil},
		attack:     {NextState: attacking, Action: nil},
		takeDamage: {NextState: dead, Action: nil},
	},
	moving: map[fsm.Input]fsm.Output{
		stop:       {NextState: idle, Action: nil},
		move:       {NextState: moving, Action: nil},
		attack:     {NextState: attacking, Action: nil},
		takeDamage: {NextState: dead, Action: nil},
	},
	attacking: map[fsm.Input]fsm.Output{
		stop:       {NextState: idle, Action: nil},
		move:       {NextState: moving, Action: nil},
		attack:     {NextState: attacking, Action: nil},
		takeDamage: {NextState: dead, Action: nil},
	},
	dead: map[fsm.Input]fsm.Output{
		stop:       {NextState: dead, Action: nil},
		move:       {NextState: dead, Action: nil},
		attack:     {NextState: dead, Action: nil},
		takeDamage: {NextState: dead, Action: nil},
	},
}

// TODO
func main() {

	fsmPool := fsm.NewFSMPool(transitionTable)

	var waitGroup sync.WaitGroup

	for i := 0; i < 200; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(50)))

		waitGroup.Add(1)

		go func(waitGroup *sync.WaitGroup, id int) {
			defer waitGroup.Done()

			fsm1 := fsmPool.Get()
			fsm1.InitState(idle)

			err := fsm1.Spin(move)

			if err != nil {
				log.Panic(err)
			}

			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(200)))
			fsmPool.Put(fsm1)
		}(&waitGroup, i)

		fmt.Println(fsmPool.Count())
	}

	waitGroup.Wait()
}
