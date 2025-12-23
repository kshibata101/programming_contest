package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

type Work struct {
	X int
	Y int
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
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
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n, d := Scan(), Scan()
	works := make([]Work, n)
	for i := 0; i < n; i++ {
		x, y := Scan(), Scan()
		works[i] = Work{x, y}
	}
	sort.Slice(works, func(i, j int) bool {
		return works[i].X < works[j].X
	})
	h := &intHeap{}
	heap.Init(h)
	j := -1
	sum := 0
	for i := 1; i <= d; i++ {
		for j < n-1 && works[j+1].X <= i {
			j++
			heap.Push(h, works[j].Y)
		}
		if h.Len() > 0 {
			sum += heap.Pop(h).(int)
		}
	}
	fmt.Println(sum)
}
