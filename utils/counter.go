package utils

import "sync/atomic"

// simple thread safe counter
type Counter interface {
	Get() int64
	Set(value int64)
	Increase() int64
	Decrease() int64
}

type counter struct {
	value int64
}

func NewCounter() Counter {
	counter := &counter{
		value: 0,
	}

	return counter
}

func (c *counter) Get() int64 {
	return atomic.LoadInt64(&c.value)
}

func (c *counter) Set(value int64) {
	atomic.StoreInt64(&c.value, value)
}

func (c *counter) Increase() int64 {
	return atomic.AddInt64(&c.value, 1)
}

func (c *counter) Decrease() int64 {
	return atomic.AddInt64(&c.value, -1)
}
