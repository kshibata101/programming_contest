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

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	s := Scan()
	dp := make([][]bool, n+1)
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, s+1)
		a[i+1] = Scan()
	}
	dp[n] = make([]bool, s+1)
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if dp[i-1][j] || (j >= a[i] && dp[i-1][j-a[i]]) {
				dp[i][j] = true
			}
		}
	}
	if dp[n][s] {
		p := []int{}
		v := s
		for i := n; i > 0; i-- {
			if dp[i-1][v] {
				// do nothing
			} else if v >= a[i] && dp[i-1][v-a[i]] {
				p = append(p, i)
				v -= a[i]
			}
		}
		fmt.Println(len(p))
		fmt.Print(p[len(p)-1])
		wr := bufio.NewWriter(os.Stdout)
		defer wr.Flush()
		for i := len(p) - 2; i >= 0; i-- {
			fmt.Fprintf(wr, " %d", p[i])
		}
		fmt.Fprintln(wr)
	} else {
		fmt.Println(-1)
	}
}
