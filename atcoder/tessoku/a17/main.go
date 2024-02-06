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
	a := make([]int, n)
	a[0] = 10000000
	for i := 1; i < n; i++ {
		a[i] = load(s)
	}
	b := make([]int, n)
	b[0] = 10000000
	b[1] = 10000000
	for i := 2; i < n; i++ {
		b[i] = load(s)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	dp := make([]int, n)
	visited := make([]bool, n)
	visited[0] = true
	visited[n-1] = true
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + a[i]
		if i >= 2 && dp[i-2]+b[i] <= dp[i] {
			dp[i] = dp[i-2] + b[i]
		}
	}
	for i := n - 1; i >= 1; {
		if i >= 2 && dp[i] == dp[i-2]+b[i] {
			visited[i-2] = true
			i -= 2
		} else {
			visited[i-1] = true
			i--
		}
	}
	k := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			k++
		}
	}
	fmt.Fprintln(w, k)
	fmt.Fprint(w, 1)
	for i := 1; i < n; i++ {
		if visited[i] {
			fmt.Fprintf(w, " %d", i+1)
		}
	}
	fmt.Fprintln(w)
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
