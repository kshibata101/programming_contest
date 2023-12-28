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
	l := 1
	for i := 0; i < n-1; i++ {
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
		// 余分に足してる分を引き算しないと
	}
	fmt.Println(sum)
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
