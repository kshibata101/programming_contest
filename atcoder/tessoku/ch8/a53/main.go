package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	q := Scan()
	pq := &intHeap{}
	heap.Init(pq)
	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			x := Scan()
			heap.Push(pq, x)
		} else if c == 2 {
			fmt.Println((*pq)[0])
		} else if c == 3 {
			heap.Pop(pq)
		} else {
			panic("invalid command")
		}
	}
}
