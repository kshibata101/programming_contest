package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

type Nd struct { // Node Distance from start point
	N int // Node
	D int // Distance
}

type NdHeap []Nd

func (h NdHeap) Len() int           { return len(h) }
func (h NdHeap) Less(i, j int) bool { return h[i].D < h[j].D }
func (h NdHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NdHeap) Push(x interface{}) {
	*h = append(*h, x.(Nd))
}
func (h *NdHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	m := Scan()
	g := make([][]Nd, n+1)
	for i := 1; i <= n; i++ {
		g[i] = []Nd{}
	}
	for i := 0; i < m; i++ {
		a := Scan()
		b := Scan()
		c := Scan()
		g[a] = append(g[a], Nd{b, c})
		g[b] = append(g[b], Nd{a, c})
	}
	pq := &NdHeap{}
	heap.Init(pq)
	heap.Push(pq, Nd{1, 0})
	dists := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dists[i] = -1
	}
	for pq.Len() > 0 {
		nd := heap.Pop(pq).(Nd)
		if dists[nd.N] != -1 {
			continue
		}

		dists[nd.N] = nd.D
		for i := 0; i < len(g[nd.N]); i++ {
			next := g[nd.N][i]
			heap.Push(pq, Nd{next.N, nd.D + next.D})
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < n; i++ {
		fmt.Fprintln(wr, dists[i+1])
	}
}
