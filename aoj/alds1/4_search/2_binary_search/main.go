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

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	S := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		S[i], _ = strconv.Atoi(s.Text())
	}

	s.Scan()
	q, _ := strconv.Atoi(s.Text())
	C := 0
	for i := 0; i < q; i++ {
		s.Scan()
		t, _ := strconv.Atoi(s.Text())

		// binary search
		l := 0
		r := n
		for l < r {
			m := (l + r) / 2
			if t < S[m] {
				r = m
			} else if t > S[m] {
				l = m + 1
			} else {
				C++
				break
			}
		}
	}
	fmt.Println(C)
}
