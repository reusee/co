package co

import (
	"math/big"
	"testing"
)

type fibState struct {
	a, b *big.Int
}

func newFibState() *fibState {
	return &fibState{
		a: big.NewInt(0),
		b: big.NewInt(1),
	}
}

func (f *fibState) next() (ret string, next Proc[string]) {
	ret = f.a.String()
	tmp := big.NewInt(0)
	tmp.Set(f.a)
	f.a.Set(f.b)
	f.b.Add(f.b, tmp)
	next = f.next
	return
}

func TestFib(t *testing.T) {
	state := newFibState()
	thread := NewThread(state, state.next)
	for {
		n, ok := thread.Step()
		if !ok {
			break
		}
		if len(n) > 1024 {
			break
		}
	}
}
