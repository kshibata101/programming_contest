package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	id      int
	l       int
	r       int
	parent  int
	depth   int
	height  int
	sibling int
}

func (x *node) degree() int {
	cnt := 0
	if x.l >= 0 {
		cnt++
	}
	if x.r >= 0 {
		cnt++
	}
	return cnt
}

func (x *node) category() string {
	if x.parent == -1 {
		return "root"
	} else if x.height == 0 {
		return "leaf"
	}
	return "internal node"
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	nodes := make([]*node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &node{id: i, parent: -1, sibling: -1}
	}
	for i := 0; i < n; i++ {
		s.Scan()
		id, _ := strconv.Atoi(s.Text())

		s.Scan()
		l, _ := strconv.Atoi(s.Text())
		nodes[id].l = l

		s.Scan()
		r, _ := strconv.Atoi(s.Text())
		nodes[id].r = r

		if l != -1 {
			nodes[l].parent = id
			nodes[l].sibling = r
		}
		if r != -1 {
			nodes[r].parent = id
			nodes[r].sibling = l
		}
	}

	rootID := -1
	for i := 0; i < n; i++ {
		if nodes[i].parent == -1 {
			rootID = nodes[i].id
			break
		}
	}
	setDepthHeight(nodes, rootID, 0)

	for i := 0; i < n; i++ {
		nd := nodes[i]
		fmt.Printf(
			"node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n",
			nd.id, nd.parent, nd.sibling, nd.degree(), nd.depth, nd.height, nd.category(),
		)
	}
}

func setDepthHeight(nodes []*node, id int, depth int) int {
	if id == -1 {
		return -1
	}
	nd := nodes[id]
	nd.depth = depth
	h1 := setDepthHeight(nodes, nd.l, depth+1)
	h2 := setDepthHeight(nodes, nd.r, depth+1)
	h := h1 + 1
	if h1 < h2 {
		h = h2 + 1
	}
	nd.height = h
	return h
}
