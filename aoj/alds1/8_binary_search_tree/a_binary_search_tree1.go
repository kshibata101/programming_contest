package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Node struct {
	Value int
	Parent *Node
	Left *Node
	Right *Node
}

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	wr := bufio.NewWriter(os.Stdout)
	
	var root *Node
	
	for i := 0; i < n; i++ {
		sc.Scan()
		command := sc.Text()
		
		if command == "print" {
			printInorder(root, wr)
			wr.WriteString("\n")
			printPreorder(root, wr)
			wr.WriteString("\n")
		} else if command == "insert" {
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			root = insert(root, k)
		}
	}
	wr.Flush()
}

func printInorder(node *Node, wr *bufio.Writer) {
	if node == nil {
		return
	}
	printInorder(node.Left, wr)
	wr.WriteString(" ")
	wr.WriteString(strconv.Itoa(node.Value))
	printInorder(node.Right, wr)
}

func printPreorder(node *Node, wr *bufio.Writer) {
	if node == nil {
		return
	}
	wr.WriteString(" ")
	wr.WriteString(strconv.Itoa(node.Value))
	printPreorder(node.Left, wr)
	printPreorder(node.Right, wr)
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
		if z.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y

	if z.Value < y.Value {
		y.Left = z
	} else {
		y.Right = z
	}
	return root
}
