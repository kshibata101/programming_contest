package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	h := make([]int, n)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = Scan()
		dp[i] = 1 << 30
	}
	dp[0] = 0

	for i := 0; i < n-1; i++ {
		c := h[i+1] - h[i]
		if c < 0 {
			c = -c
		}
		dp[i+1] = Min(dp[i+1], dp[i]+c)
		if i < n-2 {
			c = h[i+2] - h[i]
			if c < 0 {
				c = -c
			}
			dp[i+2] = Min(dp[i+2], dp[i]+c)
		}
	}
	fmt.Println(dp[n-1])
}
