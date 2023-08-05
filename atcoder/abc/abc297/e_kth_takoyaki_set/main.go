package main

import (
	"container/heap"
	"fmt"
)

type H []uint64

func (h H) Len() int           { return len(h) }
func (h H) Less(i, j int) bool { return h[i] < h[j] }
func (h H) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *H) Push(i interface{}) {
	*h = append(*h, i.(uint64))
}
func (h *H) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	a := make([]uint64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var h H = make([]uint64, n)
	copy(h, a)
	heap.Init(&h)

	c := 0
	var sum uint64 = 0
	for c < k {
		v := heap.Pop(&h).(uint64)
		if v == sum {
			continue
		}
		for i := 0; i < n; i++ {
			heap.Push(&h, v+a[i])
		}
		c++
		sum = v
	}
	fmt.Println(sum)
}
