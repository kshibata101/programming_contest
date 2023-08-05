package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n)
	fmt.Scan(&d)

	var t int
	var prev int
	fmt.Scan(&prev)
	ans := -1
	for i := 1; i < n; i++ {
		fmt.Scan(&t)
		if t-prev <= d {
			ans = t
			break
		}
		prev = t
	}
	fmt.Println(ans)
}
