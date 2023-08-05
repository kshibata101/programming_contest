package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, W int
	fmt.Scan(&N)
	fmt.Scan(&W)

	dp := make([][]int, N+1)
	dp[0] = make([]int, W+1)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for i := 0; i < N; i++ {
		s.Scan()
		w, _ := strconv.Atoi(s.Text())
		s.Scan()
		v, _ := strconv.Atoi(s.Text())

		dp[i+1] = make([]int, W+1)
		for j := 0; j <= W; j++ {
			// 加えない場合
			if dp[i][j] > dp[i+1][j] {
				dp[i+1][j] = dp[i][j]
			}
			// 加える場合
			if j+w <= W && dp[i][j]+v > dp[i+1][j+w] {
				dp[i+1][j+w] = dp[i][j] + v
			}
		}
	}

	max := 0
	for j := 0; j <= W; j++ {
		if dp[N][j] > max {
			max = dp[N][j]
		}
	}
	fmt.Println(max)
}
