package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i], _ = strconv.Atoi(sc.Text())
	}

	c := 0
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())

		l := 0
		r := n
		for l < r {
			m := (r + l) / 2
			if t == s[m] {
				c++
				break
			} else if t < s[m] {
				r = m
			} else {
				l = m + 1
			}
		}
	}
	fmt.Println(c)
}
