package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() string {
	s.Scan()
	return s.Text()
}

func ScanI() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanIArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ScanI()
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	h := ScanI()
	w := ScanI()
	mp := make([]string, h)
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		mp[i] = Scan()
		dp[i] = make([]int, w)
		if i == 0 {
			dp[i][0] = 1
		}
		for j := 0; j < w; j++ {
			if mp[i][j] == '#' {
				dp[i][j] = 0
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	fmt.Println(dp[h-1][w-1])
}
