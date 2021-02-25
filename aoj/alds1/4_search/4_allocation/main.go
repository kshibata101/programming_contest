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

	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	w := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		w[i], _ = strconv.Atoi(s.Text())
	}

	l, r := 1, 1000000000
	for l < r {
		p := (l + r) / 2
		if canCarry(w, n, k, p) {
			r = p
		} else {
			l = p + 1
		}
	}
	fmt.Println(l)
}

func canCarry(w []int, n int, k int, p int) bool {
	num, weight := 1, 0
	for i := 0; i < n; i++ {
		if w[i] > p {
			return false
		}
		if weight+w[i] > p {
			num++
			if num > k {
				return false
			}
			weight = w[i]
		} else {
			weight += w[i]
		}
	}
	return true
}
