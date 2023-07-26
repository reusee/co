package co

// Proc represents a procedure that returns a value and sets the next procedure
type Proc[V any] func(next *Proc[V]) V

// Thread represents a state and a procedure
type Thread[S any, V any] struct {
	state S
	proc  Proc[V]
}

// NewThread creates a new thread
func NewThread[S any, V any](
	initState S,
	initProc Proc[V],
) *Thread[S, V] {
	return &Thread[S, V]{
		state: initState,
		proc:  initProc,
	}
}

// Step execute the thread in single step
func (t *Thread[S, V]) Step() (v V, ok bool) {
	if t.proc == nil {
		return
	}
	v = t.proc(&t.proc)
	ok = true
	return
}
