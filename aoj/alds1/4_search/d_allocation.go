package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func check(w []int, k int, p int) bool {
	sum := 0
	for i := 0; i < len(w); i++ {
		if k <= 0 {
			return false
		}
		if sum + w[i] <= p {
			sum += w[i]
		} else {
			k--
			sum = w[i]
		}
	}
	return k > 0
}

func binarySearch(w []int, k int, l int, u int) int {
	if u - l <= 1 {
		return u
	}
	m := (u + l) / 2
	if check(w, k, m) {
		return binarySearch(w, k, l, m)
	} else {
		return binarySearch(w, k, m, u)
	}
}

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)

	pmin := 0
	psum := 0
	w := make([]int, n)
	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		s.Scan()
		w[i], _ = strconv.Atoi(s.Text())
		
		psum += w[i]
		if w[i] > pmin {
			pmin = w[i]
		}
	}

	fmt.Println(binarySearch(w, k, pmin - 1, psum))
}
