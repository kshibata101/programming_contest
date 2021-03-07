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
	t      *tree
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

func (t *tree) delete(k int) {
	n := t.root
	for n != nil {
		if k == n.k {
			break
		} else if k < n.k {
			n = n.left
		} else {
			n = n.right
		}
	}

	n.delete()
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

func (n *node) find(k int) bool {
	if n == nil {
		return false
	}
	if k < n.k {
		return n.left.find(k)
	} else if k > n.k {
		return n.right.find(k)
	} else {
		return true
	}
}

func (n *node) delete() {
	if n == nil {
		return
	}
	if n.left == nil && n.right == nil {
		if n.parent == nil {
			n.t.root = nil
		} else if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	} else if n.left == nil {
		if n.parent == nil {
			n.t.root = n.right
		} else if n.parent.left == n {
			n.parent.left = n.right
		} else {
			n.parent.right = n.right
		}
		n.right.parent = n.parent
	} else if n.right == nil {
		if n.parent == nil {
			n.t.root = n.left
		} else if n.parent.left == n {
			n.parent.left = n.left
		} else {
			n.parent.right = n.left
		}
		n.left.parent = n.parent
	} else {
		x := n.right
		for x.left != nil {
			x = x.left
		}
		n.k = x.k
		x.delete()
	}
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

			t.insert(&node{k: k, t: &t})
		} else if command == "print" {
			t.root.printInorder(w)
			fmt.Fprintln(w)
			t.root.printPreorder(w)
			fmt.Fprintln(w)
		} else if command == "find" {
			var k int
			fmt.Fscan(r, &k)

			if t.root.find(k) {
				fmt.Fprintln(w, "yes")
			} else {
				fmt.Fprintln(w, "no")
			}
		} else if command == "delete" {
			var k int
			fmt.Fscan(r, &k)

			t.delete(k)
		}
	}
}
