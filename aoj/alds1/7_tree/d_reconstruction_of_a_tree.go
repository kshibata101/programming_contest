package main

import (
	"fmt"
)

type Node struct {
	Id int
	Left *Node
	Right *Node
}

func main() {
	var n int
	fmt.Scan(&n)

	pre := make([]int, n)
	in := make([]int, n)

	for i := 0; i < n ; i++ {
		fmt.Scan(&pre[i])
	}
	for i := 0; i < n ; i++ {
		fmt.Scan(&in[i])
	}

	root := createTree(pre, in)
	queue := []int{}
	queuePostorder(&root, &queue)
	fmt.Print(queue[0])
	for i := 1; i < n; i++ {
		fmt.Printf(" %d", queue[i])
	}
	fmt.Println()
}

func createTree(pre []int, in []int) Node {
	root := Node{pre[0], nil, nil}
	
	i := 0
	leftPre := []int{}
	leftIn  := []int{}
	for ; i < len(in); i++ {
		if in[i] == root.Id {
			break
		}
		leftPre = append(leftPre, pre[i+1])
		leftIn = append(leftIn, in[i])
	}
	if len(leftPre) > 0 {
		left := createTree(leftPre, leftIn)
		root.Left = &left
	}

	i += 1
	rightPre := []int{}
	rightIn  := []int{}
	for ; i < len(in); i++ {
		rightPre = append(rightPre, pre[i])
		rightIn = append(rightIn, in[i])
	}
	if len(rightPre) > 0 {
		right := createTree(rightPre, rightIn)
		root.Right = &right
	}

	return root
}

func queuePostorder(node *Node, queue *[]int) {
	if node == nil {
		return
	}
	queuePostorder((*node).Left, queue)
	queuePostorder((*node).Right, queue)
	*queue = append(*queue, (*node).Id)
}
