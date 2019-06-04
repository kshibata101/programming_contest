package main

import (
	"fmt"
	"container/heap"
)

type Node struct {
	Id int
	Dis int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dis < pq[j].Dis
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	pop := (*pq)[n-1]
	*pq = (*pq)[0:n-1]
	return pop
}

func main() {
	var n int
	fmt.Scan(&n)

	g := make([][]int, n)
	c := make([][]int, n)
	pq := make(PriorityQueue, 0, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		var u,k int
		fmt.Scan(&u)
		fmt.Scan(&k)
		g[u] = make([]int, k)
		c[u] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&g[u][j])
			fmt.Scan(&c[u][j])
		}

		if u == 0 {
			d[u] = 0
		} else {
			d[u] = 1 << 30
		}
		pq.Push(&Node{u, d[u]})
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		for i := 0; i < len(g[node.Id]); i++ {
			nextId := g[node.Id][i]
			nextDis := node.Dis + c[node.Id][i]
			if nextDis < d[nextId] {
				d[nextId] = nextDis
				heap.Push(&pq, &Node{nextId, nextDis})
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(i, d[i])
	}
}
