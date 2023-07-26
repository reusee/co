package co

import "testing"

type counterState struct {
	n int
}

func counterNext(c *counterState) (ret Proc[int]) {
	ret = func(next *Proc[int]) int {
		c.n++
		*next = ret
		return c.n
	}
	return ret
}

func BenchmarkCounter(b *testing.B) {
	state := &counterState{n: 1}
	thread := NewThread(state, counterNext(state))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		thread.Step()
	}
}
