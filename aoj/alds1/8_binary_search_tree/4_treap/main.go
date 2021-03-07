package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	k int
	p int
	l *node
	r *node
}

func (n *node) insert(k, p int) *node {
	if n == nil {
		return &node{k: k, p: p}
	}
	if k < n.k {
		n.l = n.l.insert(k, p)
		if n.p < n.l.p {
			return n.rotateRight()
		}
	} else if k > n.k {
		n.r = n.r.insert(k, p)
		if n.p < n.r.p {
			return n.rotateLeft()
		}
	}
	return n
}

func (n *node) rotateRight() *node {
	p := n.l
	n.l = p.r
	p.r = n
	return p
}

func (n *node) rotateLeft() *node {
	p := n.r
	n.r = p.l
	p.l = n
	return p
}

func (n *node) delete(k int) *node {
	if n == nil {
		return nil
	}
	if k < n.k {
		n.l = n.l.delete(k)
	} else if k > n.k {
		n.r = n.r.delete(k)
	} else {
		return n._delete(k)
	}
	return n
}

func (n *node) _delete(k int) *node {
	if n.l == nil && n.r == nil {
		return nil
	}

	p := n
	if n.l == nil {
		p = n.rotateLeft()
	} else if n.r == nil {
		p = n.rotateRight()
	} else {
		if n.l.p < n.r.p {
			p = n.rotateLeft()
		} else {
			p = n.rotateRight()
		}
	}
	return p.delete(k)
}

func (n *node) find(k int) *node {
	if n == nil {
		return nil
	}
	if k < n.k {
		return n.l.find(k)
	} else if k > n.k {
		return n.r.find(k)
	} else {
		return n
	}
}

func (n *node) printPreorder(w io.Writer) {
	if n == nil {
		return
	}
	fmt.Fprintf(w, " %d", n.k)
	n.l.printPreorder(w)
	n.r.printPreorder(w)
}

func (n *node) printInorder(w io.Writer) {
	if n == nil {
		return
	}
	n.l.printInorder(w)
	fmt.Fprintf(w, " %d", n.k)
	n.r.printInorder(w)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	var root *node
	for i := 0; i < n; i++ {
		var command string
		fmt.Fscan(r, &command)
		if command == "insert" {
			var k, p int
			fmt.Fscan(r, &k, &p)

			root = root.insert(k, p)
		} else if command == "find" {
			var k int
			fmt.Fscan(r, &k)

			if root.find(k) == nil {
				fmt.Fprintln(w, "no")
			} else {
				fmt.Fprintln(w, "yes")
			}
		} else if command == "delete" {
			var k int
			fmt.Fscan(r, &k)

			root = root.delete(k)
		} else if command == "print" {
			root.printInorder(w)
			fmt.Fprintln(w)
			root.printPreorder(w)
			fmt.Fprintln(w)
		}
	}
}
