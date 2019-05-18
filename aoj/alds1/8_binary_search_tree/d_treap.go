package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Node struct {
	Key int
	Priority int
	Left *Node
	Right *Node
}

func main() {
	var n int
	fmt.Scan(&n)

	var root *Node
	
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		sc.Scan()
		com := sc.Text()
		if com == "print" {
			printInorder(root, wr)
			wr.WriteString("\n")
			printPreorder(root, wr)
			wr.WriteString("\n")
			continue
		}

		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		
		switch com {
		case "insert":
			sc.Scan()
			p, _ := strconv.Atoi(sc.Text())
			root = insert(root, k, p)
		case "find":
			if find(root, k) == nil {
				wr.WriteString("no\n")
			} else {
				wr.WriteString("yes\n")
			}
		case "delete":
			root = delete(root, k)
		default:
			panic("invalid command")
		}
	}
	wr.Flush()
}

func insert(t *Node, k int, p int) *Node {
	if t == nil {
		return &Node{k, p, nil, nil}
	}
	if k == t.Key {
		return t
	}

	if k < t.Key {
		t.Left = insert(t.Left, k, p)
		if t.Priority < t.Left.Priority {
			t = rotateRight(t)
		}
	} else {
		t.Right = insert(t.Right, k, p)
		if t.Priority < t.Right.Priority {
			t = rotateLeft(t)
		}
	}
	return t
}

func rotateLeft(t *Node) *Node {
	s := t.Right
	t.Right = s.Left
	s.Left = t
	return s
}

func rotateRight(t *Node) *Node {
	s := t.Left
	t.Left = s.Right
	s.Right = t
	return s
}

func find(t *Node, k int) *Node {
	if t == nil {
		return nil
	} else if t.Key == k {
		return t
	} else if k < t.Key {
		return find(t.Left, k)
	} else {
		return find(t.Right, k)
	}
}

func delete(t *Node, k int) *Node {
	if t == nil {
		return nil
	}

	if k < t.Key {
		t.Left = delete(t.Left, k)
	} else if k > t.Key {
		t.Right = delete(t.Right, k)
	} else {
		// 実際の削除
		if t.Left == nil && t.Right == nil {
			return nil
		}
		if t.Left == nil {
			t = rotateLeft(t)
		} else if t.Right == nil {
			t = rotateRight(t)
		} else {
			if t.Left.Priority < t.Right.Priority {
				t = rotateLeft(t)
			} else {
				t = rotateRight(t)
			}
		}
		return delete(t, k)
	}
	return t
}

func printInorder(t *Node, wr *bufio.Writer) {
	if t == nil {
		return
	}
	printInorder(t.Left, wr)
	wr.WriteString(" ")
	wr.WriteString(strconv.Itoa(t.Key))
	printInorder(t.Right, wr)
}

func printPreorder(t *Node, wr *bufio.Writer) {
	if t == nil {
		return
	}
	wr.WriteString(" ")
	wr.WriteString(strconv.Itoa(t.Key))
	printPreorder(t.Left, wr)
	printPreorder(t.Right, wr)
}
