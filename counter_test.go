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
	if state.n != b.N+1 {
		b.Fatalf("expected %d, got %d", b.N+1, state.n)
	}
}

func BenchmarkMultiCounter(b *testing.B) {
	state := &counterState{n: 1}
	nThreads := 1024
	var threads []*Thread[*counterState, int]
	for i := 0; i < nThreads; i++ {
		threads = append(threads, NewThread(state, counterNext(state)))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		threads[i%nThreads].Step()
	}

	if state.n != b.N+1 {
		b.Fatalf("expected %d, got %d", b.N+1, state.n)
	}
}
