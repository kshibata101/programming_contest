package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	n := load(s)
	x := load(s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = load(s)
	}

	l := 0
	r := n
	for l < r {
		m := (l + r) / 2
		if a[m] == x {
			fmt.Println(m + 1)
			return
		} else if a[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
