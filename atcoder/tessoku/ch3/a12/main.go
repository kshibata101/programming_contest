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
	l := 1
	r := int(10e9)
	for l < r {
		m := (l + r) / 2
		sum := 0
		psum := 0
		for i := 0; i < n; i++ {
			if m > 1 {
				psum += (m - 1) / a[i]
			}
			sum += m / a[i]
		}
		if sum < k {
			l = m + 1
		} else if psum >= k {
			r = m
		} else {
			fmt.Println(m)
			return
		}
	}
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
