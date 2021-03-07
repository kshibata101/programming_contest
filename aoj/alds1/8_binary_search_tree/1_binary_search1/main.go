package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	k      int
	parent *node
	left   *node
	right  *node
}

type tree struct {
	root *node
}

func (t *tree) insert(n *node) {
	var y *node
	x := t.root
	for x != nil {
		y = x
		if n.k < x.k {
			x = x.left
		} else {
			x = x.right
		}
	}
	n.parent = y

	if y == nil {
		t.root = n
	} else if n.k < y.k {
		y.left = n
	} else {
		y.right = n
	}
}

func (n *node) printInorder(w io.Writer) {
	if n == nil {
		return
	}
	n.left.printInorder(w)
	fmt.Fprintf(w, " %d", n.k)
	n.right.printInorder(w)
}

func (n *node) printPreorder(w io.Writer) {
	if n == nil {
		return
	}
	fmt.Fprintf(w, " %d", n.k)
	n.left.printPreorder(w)
	n.right.printPreorder(w)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	t := tree{}
	for i := 0; i < n; i++ {
		var command string
		fmt.Fscan(r, &command)

		if command == "insert" {
			var k int
			fmt.Fscan(r, &k)

			t.insert(&node{k: k})
		} else if command == "print" {
			t.root.printInorder(w)
			fmt.Fprintln(w)
			t.root.printPreorder(w)
			fmt.Fprintln(w)
		}
	}
}
