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
	dp[0] = make([]int, 100001)
	for j := 0; j < 100000; j++ {
		dp[0][j+1] = W + 1
	}

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for i := 0; i < N; i++ {
		s.Scan()
		w, _ := strconv.Atoi(s.Text())
		s.Scan()
		v, _ := strconv.Atoi(s.Text())

		dp[i+1] = make([]int, 100001)

		for j := 0; j <= 100000; j++ {
			if j-v >= 0 && dp[i][j-v]+w < dp[i][j] {
				dp[i+1][j] = dp[i][j-v] + w
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	maxV := 0
	for j := 0; j <= 100000; j++ {
		if dp[N][j] <= W && j > maxV {
			maxV = j
		}
	}
	fmt.Println(maxV)
}
