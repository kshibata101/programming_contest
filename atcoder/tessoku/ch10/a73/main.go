package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

type X struct {
	node int
	dist int
}

type xHeap []X

func (h xHeap) Len() int           { return len(h) }
func (h xHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h xHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *xHeap) Push(x interface{}) {
	*h = append(*h, x.(X))
}
func (h *xHeap) Pop() interface{} {
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
	xs := make([][]X, n)
	for i := 0; i < n; i++ {
		xs[i] = []X{}
	}
	for i := 0; i < m; i++ {
		a := Scan() - 1
		b := Scan() - 1
		c := Scan()
		d := Scan()
		xs[a] = append(xs[a], X{b, c*10000 - d})
		xs[b] = append(xs[b], X{a, c*10000 - d})
	}
	course := make([]X, n)
	for i := 0; i < n; i++ {
		course[i] = X{i, 1 << 31}
	}
	xh := &xHeap{}
	heap.Init(xh)
	heap.Push(xh, X{0, 0})
	for xh.Len() > 0 {
		x := heap.Pop(xh).(X)
		if x.dist < course[x.node].dist {
			// 更新
			course[x.node] = x
			for i := 0; i < len(xs[x.node]); i++ {
				next := xs[x.node][i]
				heap.Push(xh, X{next.node, x.dist + next.dist})
			}
		}
	}
	goal := course[n-1]
	dis := int(math.Ceil(float64(goal.dist) / 10000.0))
	tree := dis*10000 - goal.dist
	fmt.Printf("%d %d\n", dis, tree)
}
