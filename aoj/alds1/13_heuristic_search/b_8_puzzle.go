package main

import (
	"fmt"
	"container/heap"
)

type State struct {
	Puzzle []int
	Zero int
	Prev int
	Turn int
	Dis int
	Cost int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	pop := (*pq)[n-1]
	*pq = (*pq)[0:n-1]
	return pop
}

func main() {
	n := 9
	puzzle := make([]int, n)
	var zero int
	for i := 0; i < n; i++ {
		fmt.Scan(&puzzle[i])
		if puzzle[i] == 0 {
			puzzle[i] = n
			zero = i
		}
	}

	turn := 0
	dis := dist(&puzzle)
	if dis == 0 {
		fmt.Println(turn)
		return
	}
	que := PriorityQueue{&State{puzzle, zero, -1, turn, dis, turn + dis}}
	heap.Init(&que)
	dir := []int{-3,-1,1,3}
	for que.Len() > 0 {
		pop := heap.Pop(&que).(*State)
		for i := 0; i < len(dir); i++ {
			next := (*pop).Zero + dir[i]
			if next == (*pop).Prev {
				continue
			}
			if next < 0 || n <= next {
				continue
			}
			if (next%3 - (*pop).Zero%3) * dir[i] < 0 {
				continue
			}
			
			p := make([]int, n)
			copy(p, (*pop).Puzzle)
			p[next], p[(*pop).Zero] = p[(*pop).Zero], p[next]
			turn = (*pop).Turn + 1
			dis = dist(&p)
			if dis == 0 {
				fmt.Println(turn)
				return
			}
			
			heap.Push(&que, &State{p, next, (*pop).Zero, turn, dis, turn + dis})
		}
	}
}

func dist(p *[]int) int {
	sum := 0
	n := len(*p)
	for i := 0; i < n; i++ {
		v := (*p)[i] - 1
		sum += abs(v/3 - i/3)
		sum += abs(v%3 - i%3)
	}
	return sum / 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
