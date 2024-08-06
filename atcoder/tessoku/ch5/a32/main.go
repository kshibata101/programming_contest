package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	dp := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		if (i >= a && !dp[i-a]) || (i >= b && !dp[i-b]) {
			dp[i] = true
		}
	}
	if dp[n] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
