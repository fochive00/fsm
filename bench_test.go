package fsm_test

import "testing"

type Getter interface {
}

type getter struct {
	V int
}

func (g *getter) GetV() int {
	return g.V
}

func (g *getter) SetV(val int) {
	g.V = val
}

func Benchmark_inline(b *testing.B) {
	g := &getter{V: 1}

	for n := 0; n < b.N; n++ {
		g.SetV(5)
	}
}

func Benchmark_Direct(b *testing.B) {
	g := &getter{V: 1}

	for n := 0; n < b.N; n++ {
		g.V = 5
	}
}
