package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type P struct {
	Stage int
	Sec int
}

type pHeap []P
func (h pHeap) Len() int           { return len(h) }
func (h pHeap) Less(i, j int) bool { return h[i].Sec < h[j].Sec }
func (h pHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pHeap) Push(x interface{}) {
	*h = append(*h, x.(P))
}
func (h *pHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	n := load(s)
	a := make([]int, n)
	b := make([]int, n)
	x := make([]int, n)
	for i := 1; i < n; i++ {
		a[i] = load(s)
		b[i] = load(s)
		x[i] = load(s)
	}
	sec := make([]int, n+1)
	for i := 0; i <= n; i++ {
		sec[i] = 1 << 62
	}
	ph := &pHeap{}
	heap.Init(ph)
	sec[1] = 0
	heap.Push(ph, P{1, 0})
	for len(*ph) > 0 {
		p := heap.Pop(ph).(P)
		if p.Stage >= n || p.Sec > sec[p.Stage] {
			continue
		}
		next := p.Stage + 1
		if sec[next] > sec[p.Stage] + a[p.Stage] {
			sec[next] = sec[p.Stage] + a[p.Stage]
			heap.Push(ph, P{next, sec[next]})
		}
		next = x[p.Stage]
		if sec[next] > sec[p.Stage] + b[p.Stage] {
			sec[next] = sec[p.Stage] + b[p.Stage]
			heap.Push(ph, P{next, sec[next]})
		}
	}
	fmt.Println(sec[n])
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}