package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Problem struct {
	T int
	D int
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	problems := make([]Problem, n)
	maxD := 0
	for i := 0; i < n; i++ {
		problems[i] = Problem{Scan(), Scan()}
		if problems[i].D > maxD {
			maxD = problems[i].D
		}
	}
	sort.Slice(problems, func(i, j int) bool { return problems[i].D < problems[j].D })

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxD+1)
		for j := 0; j <= maxD; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= maxD; j++ {
			dp[i+1][j] = Max(dp[i+1][j], dp[i][j])
			if j+problems[i].T <= problems[i].D {
				dp[i+1][j+problems[i].T] = Max(dp[i+1][j+problems[i].T], dp[i][j]+1)
			}
		}
	}
	ans := 0
	for i := 0; i <= maxD; i++ {
		if dp[n][i] > ans {
			ans = dp[n][i]
		}
	}
	fmt.Println(ans)
}
