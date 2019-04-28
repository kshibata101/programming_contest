package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	fmt.Scan(&q)
	cnt := 0
	for i := 0; i < q; i++ {
		var t int
		fmt.Scan(&t)
		for j := 0; j < n; j++ {
			if t == s[j] {
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}
