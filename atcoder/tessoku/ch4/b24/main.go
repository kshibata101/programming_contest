package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

type Box struct {
	x int
	y int
}

func LowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		m := (l + r) / 2
		if arr[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	boxes := make([]Box, n)
	for i := 0; i < n; i++ {
		boxes[i] = Box{Scan(), Scan()}
	}
	// x昇順、y降順でソート
	sort.Slice(boxes, func(i, j int) bool {
		if boxes[i].x == boxes[j].x {
			return boxes[i].y > boxes[j].y
		} else {
			return boxes[i].x < boxes[j].x
		}
	})

	lis := make([]int, 0, n)
	for i := 0; i < n; i++ {
		y := boxes[i].y
		idx := LowerBound(lis, y)
		if idx >= len(lis) {
			lis = append(lis, y)
		} else {
			lis[idx] = y
		}
	}

	fmt.Println(len(lis))
}
