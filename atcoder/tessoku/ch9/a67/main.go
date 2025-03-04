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

type UnionFindTree struct {
	parent []int
	size   []int
}

func (uft *UnionFindTree) root(v int) int {
	for {
		if uft.parent[v] == -1 {
			return v
		}
		v = uft.parent[v]
	}
}

func (uft *UnionFindTree) unite(v int, u int) {
	rV := uft.root(v)
	rU := uft.root(u)
	if rV == rU {
		return
	}

	if uft.size[rV] < uft.size[rU] {
		uft.parent[rV] = rU
		uft.size[rU] += uft.size[rV]
	} else {
		uft.parent[rU] = rV
		uft.size[rV] += uft.size[rU]
	}
}

func (uft *UnionFindTree) same(v int, u int) bool {
	return uft.root(v) == uft.root(u)
}

func NewUnionFindTree(n int) *UnionFindTree {
	parent := make([]int, n+1)
	size := make([]int, n+1)
	for j := 0; j < n; j++ {
		parent[j+1] = -1
		size[j+1] = 1
	}
	return &UnionFindTree{parent: parent, size: size}
}

type Edge struct {
	a int
	b int
	c int
}

type edgeHeap []Edge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return h[i].c < h[j].c }
func (h edgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *edgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}
func (h *edgeHeap) Pop() interface{} {
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

	uft := NewUnionFindTree(n)
	eh := make(edgeHeap, m)
	for j := 0; j < m; j++ {
		a := Scan()
		b := Scan()
		c := Scan()
		eh[j] = Edge{a, b, c}
	}
	heap.Init(&eh)
	ans := 0
	for eh.Len() > 0 {
		e := heap.Pop(&eh).(Edge)
		if uft.same(e.a, e.b) {
			continue
		}
		uft.unite(e.a, e.b)
		ans += e.c
	}
	fmt.Println(ans)
}
