package fibheap

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"testing/quick"

	"github.com/google/go-cmp/cmp"
)

type operation struct {
	Push bool // if true, V is pushed
	V    int
}

func (op operation) String() string {
	if op.Push {
		return fmt.Sprintf("push %d", op.V)
	}
	return "pop"
}

type operations []operation

func (operations) Generate(rand *rand.Rand, size int) reflect.Value {
	ops := make(operations, 0, size)
	n := 0
	for i := 0; i < cap(ops); i++ {
		if n > 0 && rand.Intn(10) <= 5 {
			ops = append(ops, operation{Push: false})
			n--
		} else {
			ops = append(ops, operation{
				Push: true,
				V:    rand.Int(),
			})
			n++
		}
	}
	return reflect.ValueOf(ops)
}

type state struct {
	Len int
}

type result struct {
	States []state
	Popped []int
}

func TestProperties(t *testing.T) {
	usingSlice := func(ops operations) result {
		var r result
		r.States = append(r.States, state{
			Len: 0,
		})

		var slice []int
		for _, op := range ops {
			if op.Push {
				slice = append(slice, op.V)
				sort.Ints(slice)
			} else {
				r.Popped = append(r.Popped, slice[0])
				slice = slice[1:]
			}
			r.States = append(r.States, state{
				Len: len(slice),
			})
		}
		return r
	}
	usingHeap := func(ops operations) result {
		var r result
		r.States = append(r.States, state{
			Len: 0,
		})

		h := New[int]()
		for _, op := range ops {
			if op.Push {
				h.Push(op.V)
			} else {
				r.Popped = append(r.Popped, h.Pop())
			}
			r.States = append(r.States, state{
				Len: h.Len(),
			})
		}
		return r
	}
	if err := quick.CheckEqual(usingSlice, usingHeap, nil); err != nil {
		e := err.(*quick.CheckEqualError)
		ops := e.In[0].(operations)
		t.Errorf("ops: %v", ops)
		got := e.Out1[0].(result)
		want := e.Out2[0].(result)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("diff (-want +got)\n%s", diff)
		}
	}
}
