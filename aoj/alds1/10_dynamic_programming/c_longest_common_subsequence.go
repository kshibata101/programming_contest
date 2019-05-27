package main

import (
	"fmt"
)

func main() {
	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var x, y string
		fmt.Scan(&x)
		fmt.Scan(&y)

		xs := []rune(x)
		ys := []rune(y)

		xn := len(xs)
		yn := len(ys)
		dp := make([][]int, xn+1)
		for j := 0; j <= xn; j++ {
			dp[j] = make([]int, yn+1)
		}

		for j := 1; j <= xn; j++ {
			for k := 1; k <= yn; k++ {
				dp[j][k] = max(dp[j-1][k], dp[j][k-1])
				if xs[j-1] == ys[k-1] {
					dp[j][k] = max(dp[j][k], dp[j-1][k-1] + 1)
				} else {
					dp[j][k] = max(dp[j][k], dp[j-1][k-1])
				}
			}
		}
		fmt.Println(dp[xn][yn])
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
