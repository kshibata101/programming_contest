package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var S *bufio.Scanner

type RMQSegmentTree struct {
	node []int
	length int
}

func (t *RMQSegmentTree) Update(i int, v int) {
	i += t.length - 1
	t.node[i] = v
	for i > 0 {
		i = (i - 1) / 2
		t.node[i] = Min(t.node[2*i+1], t.node[2*i+2])
	}
}

func (t *RMQSegmentTree) RangeMin(l, r int) int {
	return t.Min(l, r, 0, 0, t.length)
}

func (t *RMQSegmentTree) Min(l, r, i, nowL, nowR int) int {
	if r <= nowL || nowR <= l {
		return INF
	} else if l <= nowL && nowR <= r {
		return t.node[i]
	}
	vl := t.Min(l, r, i*2+1, nowL, (nowL+nowR)/2)
	vr := t.Min(l, r, i*2+2, (nowL+nowR)/2, nowR)
	return Min(vl, vr)
}

func NewRMQSegmentTree(size, initV int) RMQSegmentTree {
	length := 1
	for length < size {
		length *= 2
	}
	t := RMQSegmentTree{make([]int, 2*length-1), length}
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

const INF = 1 << 31 - 1

func main() {
	S = bufio.NewScanner(os.Stdin)
	S.Split(bufio.ScanWords)
	n := ScanI()

	t := NewRMQSegmentTree(n, INF)
	q := ScanI()
	for i := 0; i < q; i++ {
		c := ScanI()
		x := ScanI()
		y := ScanI()
		if c == 0 {
			t.Update(x, y)
		} else if c == 1 {
			fmt.Println(t.RangeMin(x, y+1))
		} else {
			panic("unexpected command")
		}
	}
}

