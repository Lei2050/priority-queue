package priorityqueue

import (
	"math/rand"
	"testing"
)

func TestPQBase(t *testing.T) {
	pq := NewBasicTypePQ[int](4)
	pq.Push(1)
	pq.Push(5)
	pq.Push(9)
	pq.Push(3)
	for pq.Len() > 0 {
		t.Logf("head: %v", pq.Pop())
	}

	for i := 0; i < 20; i++ {
		pq.Push(rand.Intn(100))
	}

	for pq.Len() > 0 {
		t.Logf("head: %v", pq.Pop())
	}
}

var _ Lesser = PQStruct{}

type PQStruct struct {
	Key int
	Val int
}

// Less implements Lesser.
func (p PQStruct) Less(other Lesser) bool {
	return p.Val < other.(PQStruct).Val
}

func TestPQ(t *testing.T) {
	pq := New[PQStruct](4)
	pq.Push(PQStruct{1, 1})
	pq.Push(PQStruct{5, 5})
	pq.Push(PQStruct{9, 9})
	pq.Push(PQStruct{3, 3})

	for pq.Len() > 0 {
		t.Logf("head: %v", pq.Pop())
	}

	for i := 0; i < 20; i++ {
		pq.Push(PQStruct{rand.Intn(100), rand.Intn(100)})
	}

	for pq.Len() > 0 {
		t.Logf("head: %v", pq.Pop())
	}
}
