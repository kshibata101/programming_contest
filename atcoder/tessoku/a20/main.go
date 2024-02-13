package main

import (
	"fmt"
)

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)

	dp := make([][]int, len(S)+1)
	dp[0] = make([]int, len(T)+1)
	for i := 0; i < len(S); i++ {
		dp[i+1] = make([]int, len(T)+1)
		for j := 0; j < len(T); j++ {
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			if string(S[i]) == string(T[j]) {
				dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]+1)
			}
		}
	}
	fmt.Println(dp[len(S)][len(T)])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
