package main

import (
	"bufio"
	"os"
	"strconv"
)

var S *bufio.Scanner

type RSQSegmentTree struct {
	node []int
	length int
}

func (t *RSQSegmentTree) Add(i, v int) {
	i += t.length - 1
	t.node[i] += v
	for i > 0 {
		i = (i - 1) / 2
		t.node[i] = t.node[2*i+1] + t.node[2*i+2]
	}
}

func (t *RSQSegmentTree) Update(i int, v int) {
	i += t.length - 1
	t.node[i] = v
	for i > 0 {
		i = (i - 1) / 2
		t.node[i] = t.node[2*i+1] + t.node[2*i+2]
	}
}

func (t *RSQSegmentTree) RangeSum(l, r int) int {
	return t.Sum(l, r, 0, 0, t.length)
}

func (t *RSQSegmentTree) Sum(l, r, i, nowL, nowR int) int {
	if r <= nowL || nowR <= l {
		return 0
	} else if l <= nowL && nowR <= r {
		return t.node[i]
	}
	vl := t.Sum(l, r, i*2+1, nowL, (nowL+nowR)/2)
	vr := t.Sum(l, r, i*2+2, (nowL+nowR)/2, nowR)
	return vl + vr
}

func NewRSQSegmentTree(size, initV int) RSQSegmentTree {
	length := 1
	for length < size {
		length *= 2
	}
	t := RSQSegmentTree{make([]int, 2*length-1), length}
	for i := range t.node {
		t.node[i] = initV
	}
	return t
}

func ScanI() int {
	S.Scan()
	a, _ := strconv.Atoi(S.Text())
	return a
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	S = bufio.NewScanner(os.Stdin)
	S.Split(bufio.ScanWords)
}

