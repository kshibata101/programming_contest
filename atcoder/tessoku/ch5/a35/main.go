package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
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
	n := Scan()
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[n-1] = ScanArr(n)
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i%2 == 0 { // 太郎
				dp[i][j] = Max(dp[i+1][j], dp[i+1][j+1])
			} else { // 次郎
				dp[i][j] = Min(dp[i+1][j], dp[i+1][j+1])
			}
		}
	}
	fmt.Println(dp[0][0])
}
