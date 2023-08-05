package main

import (
	"fmt"
)

func main() {
	var n, m int64
	fmt.Scan(&n)
	fmt.Scan(&m)

	var ans int64 = 1 << 62
	var a int64
	for a = 1; a <= n; a++ {
		b := (m + a - 1) / a
		if b < a {
			break
		}
		if b <= n && a * b < ans {
			ans = a * b
		}
	}
	if ans == 1 << 62 {
		ans = -1
	}
	fmt.Println(ans)
}