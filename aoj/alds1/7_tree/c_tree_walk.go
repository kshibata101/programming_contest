package main

import (
	"fmt"
)

type Node struct {
	Id int
	Left int
	Right int
	Parent int
}

func main() {
	var n int
	fmt.Scan(&n)

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].Parent = -1
	}
	
	for i := 0; i < n; i++ {
		var id int
		fmt.Scan(&id)
		nodes[id].Id = id
		fmt.Scan(&nodes[id].Left)
		fmt.Scan(&nodes[id].Right)

		if nodes[id].Left >= 0 {
			nodes[nodes[id].Left].Parent = nodes[id].Id
		}
		if nodes[id].Right >= 0 {
			nodes[nodes[id].Right].Parent = nodes[id].Id
		}
	}

	parent := -1
	for i := 0; i < n; i++ {
		if nodes[i].Parent == -1 {
			parent = nodes[i].Id
		}
	}
	fmt.Println("Preorder")
	printPreorder(&nodes, parent)
	fmt.Println()

	fmt.Println("Inorder")
	printInorder(&nodes, parent)
	fmt.Println()
	
	fmt.Println("Postorder")
	printPostorder(&nodes, parent)
	fmt.Println()
}

func printPreorder(nodes *[]Node, id int) {
	if id < 0 {
		return
	}
	fmt.Printf(" %d", id)
	printPreorder(nodes, (*nodes)[id].Left)
	printPreorder(nodes, (*nodes)[id].Right)
}

func printInorder(nodes *[]Node, id int) {
	if id < 0 {
		return
	}
	printInorder(nodes, (*nodes)[id].Left)
	fmt.Printf(" %d", id)
	printInorder(nodes, (*nodes)[id].Right)
}

func printPostorder(nodes *[]Node, id int) {
	if id < 0 {
		return
	}
	printPostorder(nodes, (*nodes)[id].Left)
	printPostorder(nodes, (*nodes)[id].Right)
	fmt.Printf(" %d", id)
}
