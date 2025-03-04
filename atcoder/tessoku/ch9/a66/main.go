package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

type UnionFindTree struct {
	parent []int
	size   []int
}

func (uft *UnionFindTree) root(v int) int {
	for {
		if uft.parent[v] == -1 {
			return v
		}
		v = uft.parent[v]
	}
}

func (uft *UnionFindTree) same(u int, v int) bool {
	return uft.root(u) == uft.root(v)
}

func (uft *UnionFindTree) unite(u int, v int) {
	rU := uft.root(u)
	rV := uft.root(v)
	if rU == rV {
		return
	}
	sU := uft.size[rU]
	sV := uft.size[rV]
	if sU < sV {
		uft.parent[rU] = rV
		uft.size[rV] += uft.size[rU]
	} else {
		uft.parent[rV] = rU
		uft.size[rU] += uft.size[rV]
	}
}

func NewUnionFindTree(n int) *UnionFindTree {
	parent := make([]int, n+1)
	size := make([]int, n+1)
	for j := 0; j < n; j++ {
		parent[j+1] = -1
		size[j+1] = 1
	}
	return &UnionFindTree{parent, size}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	uft := NewUnionFindTree(n)
	q := Scan()
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i := 0; i < q; i++ {
		c := Scan()
		u := Scan()
		v := Scan()
		if c == 1 {
			uft.unite(u, v)
		} else if c == 2 {
			if uft.same(u, v) {
				fmt.Fprintln(wr, "Yes")
			} else {
				fmt.Fprintln(wr, "No")
			}
		} else {
			panic("")
		}
	}
}
