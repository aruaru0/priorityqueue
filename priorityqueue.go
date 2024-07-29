package priorityqueue

import (
	"cmp"
)

// Queue ... 実装例
type Queue[T cmp.Ordered] []T

func (q Queue[T]) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q Queue[T]) Len() int {
	return len(q)
}

func (q Queue[T]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue[T]) Push(x T) {
	*q = append(*q, x)
	cur := q.Len()
	parent := cur / 2
	for cur != 1 {
		if q.Less(cur-1, parent-1) {
			q.Swap(cur-1, parent-1)
		} else {
			break
		}
		cur = parent
		parent = cur / 2
	}
}

func (q *Queue[T]) Pop() T {
	if q.Len() == 0 {
		panic("queue is empty")
	}
	old := *q
	n := len(old)
	item := old[0]

	old[0] = old[n-1]
	old = old[:n-1]
	cur := 1
	for {
		nxt0 := cur * 2
		nxt1 := cur*2 + 1
		if nxt0 > len(old) {
			break
		}
		nxt := nxt0
		if nxt1 <= len(old) && old.Less(nxt1-1, nxt0-1) {
			nxt = nxt1
		}
		if old.Less(nxt-1, cur-1) {
			old.Swap(nxt-1, cur-1)
		} else {
			break
		}

		cur = nxt
	}

	*q = old
	return item
}
