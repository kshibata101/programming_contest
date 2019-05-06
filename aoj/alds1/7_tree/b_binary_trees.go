package main

import (
	"fmt"
)

type Node struct {
	Id int
	Left int
	Right int
	Parent int
	Depth int
	Height int
	Slibing int
	Degree int
}

func main() {
	var n int
	fmt.Scan(&n)

	// init
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].Parent = -1
		nodes[i].Depth = -1
		nodes[i].Left = -1
		nodes[i].Right = -1
		nodes[i].Slibing = -1
	}

	// input
	for i := 0; i < n; i++ {
		var id int
		fmt.Scan(&id)
		nodes[id].Id = id
		fmt.Scan(&(nodes[id].Left))
		fmt.Scan(&(nodes[id].Right))
		if nodes[id].Left >= 0 {
			nodes[nodes[id].Left].Parent = nodes[id].Id
		}
		if nodes[id].Right >= 0 {
			nodes[nodes[id].Right].Parent = nodes[id].Id
		}
	}

	// find parent
	parent := -1
	for i := 0; i < n; i++ {
		if nodes[i].Parent == -1 {
			parent = nodes[i].Id
			break
		}
	}

	if parent < 0 {
		panic("invalid parent")
	}
	find(&nodes, parent, 0, -1)

	for i := 0; i < n; i++ {
		node := nodes[i]
		ntype := "leaf"
		if node.Parent == -1 {
			ntype = "root"
		} else if node.Degree > 0 {
			ntype = "internal node"
		}
		fmt.Printf("node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n", node.Id, node.Parent, node.Slibing, node.Degree, node.Depth, node.Height, ntype)
	}
}

func find(nodes *[]Node, id int, depth int, slibing int) {
	node := &(*nodes)[id]
	node.Depth = depth
	node.Slibing = slibing
	if node.Left >= 0 {
		node.Degree++
		find(nodes, node.Left, depth + 1, node.Right)
		if node.Height < (*nodes)[node.Left].Height + 1 {
			node.Height = (*nodes)[node.Left].Height + 1
		}
	}
	if node.Right >= 0 {
		node.Degree++
		find(nodes, node.Right, depth + 1, node.Left)
		if node.Height < (*nodes)[node.Right].Height + 1 {
			node.Height = (*nodes)[node.Right].Height + 1
		}
	}
}

