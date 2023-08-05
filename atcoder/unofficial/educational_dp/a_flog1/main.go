package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	dp := make([]int, n)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		s.Scan()
		h[i], _ = strconv.Atoi(s.Text())
		dp[i] = 1 << 20
	}

	dp[0] = 0
	dp[1] = abs(h[0], h[1])
	for j := 2; j < n; j++ {
		s1 := dp[j-2] + abs(h[j], h[j-2])
		s2 := dp[j-1] + abs(h[j], h[j-1])
		if s1 < s2 {
			dp[j] = s1
		} else {
			dp[j] = s2
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
