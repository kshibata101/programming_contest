package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	id     int
	l      int
	r      int
	parent int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	nodes := make([]*node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &node{
			id:     i,
			parent: -1,
		}
	}
	for i := 0; i < n; i++ {
		s.Scan()
		id, _ := strconv.Atoi(s.Text())

		s.Scan()
		l, _ := strconv.Atoi(s.Text())
		nodes[id].l = l
		if l != -1 {
			nodes[l].parent = id
		}

		s.Scan()
		r, _ := strconv.Atoi(s.Text())
		nodes[id].r = r
		if r != -1 {
			nodes[r].parent = id
		}
	}
	var rootID int
	for i := 0; i < n; i++ {
		if nodes[i].parent == -1 {
			rootID = nodes[i].id
		}
	}

	fmt.Println("Preorder")
	preorder(nodes, rootID)
	fmt.Println()

	fmt.Println("Inorder")
	inorder(nodes, rootID)
	fmt.Println()

	fmt.Println("Postorder")
	postorder(nodes, rootID)
	fmt.Println()
}

func preorder(nodes []*node, id int) {
	if id == -1 {
		return
	}
	fmt.Printf(" %d", id)
	preorder(nodes, nodes[id].l)
	preorder(nodes, nodes[id].r)
}

func inorder(nodes []*node, id int) {
	if id == -1 {
		return
	}
	inorder(nodes, nodes[id].l)
	fmt.Printf(" %d", id)
	inorder(nodes, nodes[id].r)
}

func postorder(nodes []*node, id int) {
	if id == -1 {
		return
	}
	postorder(nodes, nodes[id].l)
	postorder(nodes, nodes[id].r)
	fmt.Printf(" %d", id)
}
