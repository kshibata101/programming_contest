package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	id       int
	parent   int
	depth    int
	children []int
}

func (x *node) category() string {
	if x.parent == -1 {
		return "root"
	} else if len(x.children) > 0 {
		return "internal node"
	} else {
		return "leaf"
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	nodes := make([]*node, n) // index = order
	tree := make([]*node, n)  // index = id
	for i := 0; i < n; i++ {
		nd := node{
			parent: -1,
		}
		s.Scan()
		nd.id, _ = strconv.Atoi(s.Text())

		s.Scan()
		k, _ := strconv.Atoi(s.Text())
		nd.children = make([]int, k)

		for j := 0; j < k; j++ {
			s.Scan()
			nd.children[j], _ = strconv.Atoi(s.Text())
		}
		nodes[i] = &nd
		tree[nd.id] = &nd
	}

	for i := 0; i < n; i++ {
		nd := nodes[i]
		for j := 0; j < len(nd.children); j++ {
			tree[nd.children[j]].parent = nd.id
		}
	}
	rootID := -1
	for i := 0; i < n; i++ {
		if nodes[i].parent == -1 {
			rootID = nodes[i].id
			break
		}
	}
	setDepth(tree, rootID, 0)

	for i := 0; i < n; i++ {
		x := tree[i]
		fmt.Printf(
			"node %d: parent = %d, depth = %d, %s, [",
			x.id, x.parent, x.depth, x.category(),
		)
		if len(x.children) > 0 {
			fmt.Print(x.children[0])
			for j := 1; j < len(x.children); j++ {
				fmt.Printf(", %d", x.children[j])
			}
		}
		fmt.Println("]")
	}
}

func setDepth(tree []*node, id, depth int) {
	tree[id].depth = depth
	for i := 0; i < len(tree[id].children); i++ {
		setDepth(tree, tree[id].children[i], depth+1)
	}
}
