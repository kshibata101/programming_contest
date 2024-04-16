package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := load(s)
	S := load(s)
	a := loadArr(s, n)

	dp := make([][]bool, n+1)
	dp[0] = make([]bool, S+1)
	dp[0][0] = true
	for i := 0; i < n; i++ {
		dp[i+1] = make([]bool, S+1)
		for j := 0; j <= S; j++ {
			if dp[i][j] {
				dp[i+1][j] = true
				if j+a[i] <= S {
					dp[i+1][j+a[i]] = true
				}
			}
		}
	}
	if dp[n][S] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func loadArr(s *bufio.Scanner, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = load(s)
	}
	return a
}
