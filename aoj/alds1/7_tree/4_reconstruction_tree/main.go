package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id int
	l  *node
	r  *node
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	in := make([]int, n)
	pre := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		pre[i], _ = strconv.Atoi(s.Text())
	}
	for i := 0; i < n; i++ {
		s.Scan()
		in[i], _ = strconv.Atoi(s.Text())
	}

	root := tree(pre, in)
	que := []int{}
	queuePostOrder(root, &que)

	fmt.Println(strings.Trim(fmt.Sprint(que), "[]"))
}

func tree(pre []int, in []int) *node {
	if len(pre) <= 0 {
		return nil
	}
	n := node{id: pre[0], l: nil, r: nil}

	i := 0
	for ; i < len(pre); i++ {
		if in[i] == pre[0] {
			break
		}
	}
	n.l = tree(pre[1:i+1], in[:i])
	if i+1 < len(pre) {
		n.r = tree(pre[i+1:], in[i+1:])
	}

	return &n
}

func queuePostOrder(n *node, arr *[]int) {
	if n == nil {
		return
	}
	queuePostOrder(n.l, arr)
	queuePostOrder(n.r, arr)
	*arr = append(*arr, n.id)
}
