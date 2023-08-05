package main

import "fmt"

func main() {
	var n, m uint64
	fmt.Scan(&n)
	fmt.Scan(&m)

	dp := make([][]uint64, n)
	dp[0] = make([]uint64, 2)
	dp[0][0] = m

	var i uint64
	for i = 1; i < n; i++ {
		dp[i] = make([]uint64, 2)
		dp[i][0] = dp[i-1][1]
		dp[i][1] = (dp[i-1][0]*(m-1) + dp[i-1][1]*(m-2)) % 998244353
	}
	fmt.Println(dp[n-1][1])
}
