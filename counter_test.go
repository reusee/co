package co

import "testing"

type counterState struct {
	n int
}

func (c *counterState) next(proc *Proc[int]) int {
	c.n++
	*proc = c.next
	return c.n
}

func BenchmarkCounter(b *testing.B) {
	state := &counterState{n: 1}
	thread := NewThread(state, state.next)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		thread.Step()
	}
}
