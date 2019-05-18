package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
	Parent *Node
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
		command := sc.Text()
		switch command {
		case "print":
			printInorder(wr, root)
			wr.WriteString("\n")
			printPreorder(wr, root)
			wr.WriteString("\n")
		case "insert":
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			root = insert(root, k)
		case "find":
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			if find(root, k) {
				wr.WriteString("yes\n")
			} else {
				wr.WriteString("no\n")
			}
			
		default:
			panic("undefined command")
		}
	}
	wr.Flush()
}

func printPreorder(wr *bufio.Writer, node *Node) {
	if node == nil {
		return
	}
	wr.WriteString(" ")
	wr.WriteString(strconv.Itoa(node.Value))
	printPreorder(wr, node.Left)
	printPreorder(wr, node.Right)
}

func printInorder(wr *bufio.Writer, node *Node) {
	if node == nil {
		return
	}
	printInorder(wr, node.Left)
	wr.WriteString(" ")
	wr.WriteString(strconv.Itoa(node.Value))
	printInorder(wr, node.Right)
}

func insert(root *Node, k int) *Node {
	z := &Node{k, nil, nil, nil}
	if root == nil {
		return z
	}
	
	var x, y *Node
	x = root
	for x != nil {
		y = x
		if k < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.Parent = y
	if k < y.Value {
		y.Left = z
	} else {
		y.Right = z
	}
	return root
}

func find(node *Node, k int) bool {
	if node == nil {
		return false
	} else if node.Value == k {
		return true
	} else if k < node.Value {
		return find(node.Left, k)
	} else {
		return find(node.Right, k)
	}
}
