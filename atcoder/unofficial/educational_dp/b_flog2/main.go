package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	h := make([]int, n)
	dp := make([]int, n)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		s.Scan()
		h[i], _ = strconv.Atoi(s.Text())
		dp[i] = 1 << 31
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 1; j <= k && i+j < n; j++ {
			d := abs(h[i], h[i+j])
			if dp[i]+d < dp[i+j] {
				dp[i+j] = dp[i] + d
			}
		}
	}

	fmt.Println(dp[n-1])
}

func abs(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
