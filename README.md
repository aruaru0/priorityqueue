# priorityQueue
## Priority queue with Generics support

* Supports only cmp.Ordered types

usage 
```go
var pq Queue[int]

pq.Push(2)
pq.Push(3)
pq.Push(1)

for len(pq) != 0 {
    e = pq.Pop()
    fmt.Println(e)
}
```

