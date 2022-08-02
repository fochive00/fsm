package fsm

import "sync"

// simple thread safe counter for benchmark
type Counter struct {
	sync.Mutex
	value uint
}

func NewCounter() *Counter {
	counter := &Counter{
		value: 0,
	}

	return counter
}

func (c *Counter) Get() uint {
	c.Lock()
	defer c.Unlock()

	return c.value
}

func (c *Counter) Set(value uint) {
	c.Lock()
	defer c.Unlock()

	c.value = value
}

func (c *Counter) Increase() {
	c.Lock()
	defer c.Unlock()

	c.value++
}

func (c *Counter) Decrease() {
	c.Lock()
	defer c.Unlock()

	c.value--
}
