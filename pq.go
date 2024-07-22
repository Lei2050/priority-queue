package priorityqueue

import (
	"container/heap"
)

type Lesser interface {
	Less(Lesser) bool
}

type priorityQueueWrap[T Lesser] []T

func (pq priorityQueueWrap[T]) Len() int {
	return len(pq)
}

func (pq priorityQueueWrap[T]) Less(i, j int) bool {
	return pq[i].Less(pq[j])
}

func (pq priorityQueueWrap[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueueWrap[T]) Push(x any) {
	tmp := *pq
	n := len(tmp)
	tmp = tmp[0 : n+1]
	elem := x.(T)
	tmp[n] = elem
	*pq = tmp
}

func (pq *priorityQueueWrap[T]) Pop() any {
	tmp := *pq
	n := len(tmp)
	x := tmp[n-1]
	*pq = tmp[0 : n-1]
	return x
}

type PriorityQueue[T Lesser] struct {
	data priorityQueueWrap[T]
}

func New[T Lesser](cap int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		data: make(priorityQueueWrap[T], 0, cap),
	}
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.data.Len()
}

func (pq *PriorityQueue[T]) Cap() int {
	return cap(pq.data)
}

func (pq *PriorityQueue[T]) Push(x T) {
	l := pq.data.Len()
	if l >= cap(pq.data) {
		pq.data = append(pq.data, x)
		pq.data = pq.data[:l]
	}
	heap.Push(&pq.data, x)
}

func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(&pq.data).(T)
}
