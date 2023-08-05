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

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	dp := make([][]int, n+1)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
		s.Scan()
		b[i], _ = strconv.Atoi(s.Text())
		s.Scan()
		c[i], _ = strconv.Atoi(s.Text())

		dp[i] = make([]int, 3)
	}
	dp[n] = make([]int, 3)

	for i := 0; i < n; i++ {
		dp[i+1][0] = max(dp[i+1][0], dp[i][1]+b[i], dp[i][2]+c[i])
		dp[i+1][1] = max(dp[i+1][1], dp[i][2]+c[i], dp[i][0]+a[i])
		dp[i+1][2] = max(dp[i+1][2], dp[i][0]+a[i], dp[i][1]+b[i])
	}
	fmt.Println(max(dp[n][0], dp[n][1], dp[n][2]))
}

func max(x, y, z int) int {
	if x > y {
		if x > z {
			return x
		} else {
			return z
		}
	} else {
		if y > z {
			return y
		} else {
			return z
		}
	}
}
