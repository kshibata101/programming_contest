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

type DP struct {
	c int
	f int
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	h := make([]int, n)
	dp := make([]DP, n)
	for i := 0; i < n; i++ {
		h[i] = Scan()
		dp[i] = DP{1 << 30, -1}
	}
	dp[0] = DP{0, -1}
	for i := 1; i < n; i++ {
		c1 := Max(h[i]-h[i-1], h[i-1]-h[i])
		if i >= 2 {
			c2 := Max(h[i]-h[i-2], h[i-2]-h[i])
			if dp[i-1].c+c1 < dp[i-2].c+c2 {
				dp[i] = DP{dp[i-1].c + c1, i - 1}
			} else {
				dp[i] = DP{dp[i-2].c + c2, i - 2}
			}
		} else {
			dp[1] = DP{dp[0].c + c1, 0}
		}
	}
	p := []int{}
	f := n - 1
	for f >= 0 {
		p = append(p, f)
		f = dp[f].f
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	fmt.Fprintln(wr, len(p))
	fmt.Fprint(wr, p[len(p)-1]+1)
	for i := len(p) - 2; i >= 0; i-- {
		fmt.Fprintf(wr, " %d", p[i]+1)
	}
	fmt.Fprintln(wr)
}
