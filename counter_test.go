package co

import "testing"

type counterState struct {
	n int
}

func (c *counterState) next() (int, Proc[int]) {
	c.n++
	return c.n, c.next
}

func BenchmarkCounter(b *testing.B) {
	state := &counterState{n: 1}
	thread := NewThread(state, state.next)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		thread.Step()
	}
}
