package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func ScanI() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanIArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ScanI()
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := ScanI()
	l := []int{}
	for i := 0; i < n; i++ {
		a := ScanI()
		ins := search(a, l)
		if ins >= len(l) {
			l = append(l, a)
		} else {
			l[ins] = a
		}
	}
	fmt.Println(len(l))
}

func search(v int, a []int) int {
	l := 0
	r := len(a)
	for l < r {
		m := (l + r) / 2
		if v < a[m] {
			r = m
		} else if v > a[m] {
			l = m + 1
		} else {
			return m
		}
	}
	return l
}
