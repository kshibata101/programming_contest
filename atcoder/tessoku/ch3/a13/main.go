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
	k := load(s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = load(s)
	}

	sum := 0
	l := 0
	for i := 0; i < n-1; i++ {
		prevl := l
		r := n - 1
		for l < r {
			m := (l + r + 1) / 2
			if a[m]-a[i] < k {
				l = m
			} else if a[m]-a[i] > k {
				r = m - 1
			} else {
				l = m
				r = m
			}
		}

		sum += (l - i + 1) * (l - i) / 2
		sum -= (prevl - i + 1) * (prevl - i) / 2
	}
	fmt.Println(sum)
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
