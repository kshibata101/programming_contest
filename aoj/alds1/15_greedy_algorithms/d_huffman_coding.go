package main

import (
	"fmt"
	"container/heap"
)

type Node struct {
	C rune
	H int
	P *Node
	C0 *Node
	C1 *Node
}

type Npq []*Node

func (npq Npq) Len() int { return len(npq) }
func (npq Npq) Swap(i,j int) { npq[i], npq[j] = npq[j], npq[i] }
func (npq Npq) Less(i,j int) bool { return npq[i].H < npq[j].H }
func (npq *Npq) Push(x interface{}) {
	*npq = append(*npq, x.(*Node))
}
func (npq *Npq) Pop() interface{} {
	n := len(*npq)
	item := (*npq)[n-1]
	*npq = (*npq)[0:n-1]
	return item
}

func main() {
	var s string
	fmt.Scan(&s)
	ss := []rune(s)
	m := make(map[rune]int)
	for i := 0; i < len(ss); i++ {
		if _, ok := m[ss[i]]; !ok {
			m[ss[i]] = 0
		}
		m[ss[i]] += 1
	}

	n := len(m)
	var npq Npq = make([]*Node, n)
	i := 0
	for k, v := range m {
		npq[i] = &Node{k, v, nil, nil, nil}
		i++
	}
	heap.Init(&npq)

	for len(npq) > 1 {
		n1 := heap.Pop(&npq).(*Node)
		n2 := heap.Pop(&npq).(*Node)
		node := &Node{C: -1, H: n1.H + n2.H, P:nil, C0: n1, C1: n2}
		n1.P = node
		n2.P = node
		heap.Push(&npq, node)
	}

	que := make([]*Node, 1, 2)
	que[0] = heap.Pop(&npq).(*Node)
	if que[0].C0 == nil && que[0].C1 == nil {
		fmt.Println(que[0].H)
		return
	}

	val, d := 0, 0
	for len(que) > 0 {
		size := len(que)
		for j := 0; j < size; j++ {
			node := que[j]
			if node.C != -1 {
				val += d * node.H
			}
			if node.C0 != nil {
				que = append(que, node.C0)
			}
			if node.C1 != nil {
				que = append(que, node.C1)
			}
		}
		if len(que) == size {
			break
		}
		que = que[size:]
		d++
	}
	fmt.Println(val)
}