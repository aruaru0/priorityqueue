package priorityqueue

import (
	"fmt"
	"testing"
)

func TestIntQueue(t *testing.T) {
	fmt.Println("TestIntQueue")
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
	fmt.Println("TestStringQueue")
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
	fmt.Println("TestEmpty")
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

type pair struct {
	x, y int
}

func TestQueueL(t *testing.T) {
	fmt.Println("TestQueueL")
	pq := New[pair](func(i, j pair) bool {
		return i.x < j.x
	})

	pq.Push(pair{2, 1})
	pq.Push(pair{3, -1})
	pq.Push(pair{1, 10})

	if pq.Len() != 3 {
		t.FailNow()
	}

	for i := 1; i <= 3; i++ {
		v, ok := pq.Pop()
		if v.x != i || ok == false {
			t.FailNow()
		}
	}

	if pq.Len() != 0 {
		t.FailNow()
	}

	_, ok := pq.Pop()
	if ok == true {
		t.FailNow()
	}
}

type datas struct {
	x, y, z int
	s       string
}

func TestQueueL2(t *testing.T) {
	fmt.Println("TestQueueL2")
	pq := New[datas](func(i, j datas) bool {
		return i.s < j.s
	})

	pq.Push(datas{2, 1, 3, "aaa"})
	pq.Push(datas{3, -1, 3, "bbb"})
	pq.Push(datas{1, 10, 4, "ccc"})
	pq.Push(datas{2, 1, 3, "xyz"})
	pq.Push(datas{3, -1, 3, "bbb"})

	if pq.Len() != 5 {
		t.FailNow()
	}

	for _, e := range []string{"aaa", "bbb", "bbb", "ccc", "xyz"} {
		v, ok := pq.Pop()
		if ok == false {
			break
		}
		if v.s != e {
			t.FailNow()
		}
	}

	if pq.Len() != 0 {
		t.FailNow()
	}

	_, ok := pq.Pop()
	if ok == true {
		t.FailNow()
	}
}
