package priorityqueue

import (
	"testing"
)

func TestIntQueue(t *testing.T) {
	pq := make(Queue[int], 0)

	pq.Push(2)
	pq.Push(3)
	pq.Push(1)

	if len(pq) != 3 {
		t.FailNow()
	}

	for i := 1; i <= 3; i++ {
		v := pq.Pop()
		if v != i {
			t.FailNow()
		}
	}
}

func TestStringQueue(t *testing.T) {
	pq := make(Queue[string], 0)

	pq.Push("xxx")
	pq.Push("abc")
	pq.Push("aaa")

	if len(pq) != 3 {
		t.FailNow()
	}

	v := pq.Pop()
	if v != "aaa" {
		t.FailNow()
	}

	v = pq.Pop()
	if v != "abc" {
		t.FailNow()
	}

	v = pq.Pop()
	if v != "xxx" {
		t.FailNow()
	}
}

func TestEmpty(t *testing.T) {
	defer func() {
		err := recover()
		if err != "queue is empty" {
			t.Errorf("got %v\nwant %v", err, "queue is empty")
		}
	}()

	var pq Queue[int]

	if pq.Len() != 0 {
		t.FailNow()
	}

	pq.Pop()
}
