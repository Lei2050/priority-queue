package priorityqueue

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type basicTypePQWrap[T constraints.Ordered] []T

func (pq basicTypePQWrap[T]) Len() int {
	return len(pq)
}

func (pq basicTypePQWrap[T]) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq basicTypePQWrap[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *basicTypePQWrap[T]) Push(x any) {
	tmp := *pq
	n := len(tmp)
	tmp = tmp[0 : n+1]
	elem := x.(T)
	tmp[n] = elem
	*pq = tmp
}

func (pq *basicTypePQWrap[T]) Pop() any {
	tmp := *pq
	n := len(tmp)
	x := tmp[n-1]
	*pq = tmp[0 : n-1]
	return x
}

type BasicTypePriorityQueue[T constraints.Ordered] struct {
	data basicTypePQWrap[T]
}

func NewBasicTypePQ[T constraints.Ordered](cap int) *BasicTypePriorityQueue[T] {
	return &BasicTypePriorityQueue[T]{
		data: make(basicTypePQWrap[T], 0, cap),
	}
}

func (pq *BasicTypePriorityQueue[T]) Len() int {
	return pq.data.Len()
}

func (pq *BasicTypePriorityQueue[T]) Cap() int {
	return cap(pq.data)
}

func (pq *BasicTypePriorityQueue[T]) Push(x T) {
	l := pq.data.Len()
	if l >= cap(pq.data) {
		pq.data = append(pq.data, x)
		pq.data = pq.data[:l]
	}
	heap.Push(&pq.data, x)
}

func (pq *BasicTypePriorityQueue[T]) Pop() T {
	return heap.Pop(&pq.data).(T)
}
