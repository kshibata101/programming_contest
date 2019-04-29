package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent int
	Children []int
	Depth int
}

func main() {
	var n int
	fmt.Scan(&n)

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].Parent = -1
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		id, _ := strconv.Atoi(sc.Text())
		
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		for j := 0; j < k; j++ {
			sc.Scan()
			c, _ := strconv.Atoi(sc.Text())
			nodes[id].Children = append(nodes[id].Children, c)
			nodes[c].Parent = id
		}
	}

	root := -1
	for i := 0; i < n; i++ {
		if nodes[i].Parent == -1 {
			root = i
			break
		}
	}
	setDepth(nodes, root, 0)

	for i := 0; i < n; i++ {
		no := nodes[i]
		t := "leaf"
		if i == root {
			t = "root"
		} else if len(no.Children) > 0 {
			t = "internal node"
		}
		children := strings.Replace(fmt.Sprintf("%v", no.Children), " ", ", ", -1)
		fmt.Printf("node %d: parent = %d, depth = %d, %s, %s\n", i, no.Parent, no.Depth, t, children)
	}
}

func setDepth(nodes []Node, current int, depth int) {
	nodes[current].Depth = depth
	for i := 0; i < len(nodes[current].Children); i++ {
		setDepth(nodes, nodes[current].Children[i], depth + 1)
	}
}
