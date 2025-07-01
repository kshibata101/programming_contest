package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func binarySearch(v, l, r int, x []int, up bool) int {
	for l < r {
		m := (l + r) / 2
		if v < x[m] {
			return binarySearch(v, l, m, x, up)
		} else if v > x[m] {
			return binarySearch(v, m+1, r, x, up)
		} else {
			return m
		}
	}
	if up {
		return l
	} else {
		return l - 1
	}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	w := Scan()
	l := Scan()
	r := Scan()
	x := make([]int, n+2)
	x[0] = 0
	for i := 1; i <= n; i++ {
		x[i] = Scan()
	}
	x[n+1] = w
	dp := make([]int, n+2)
	sum := make([]int, n+2)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		// dp[i]を考える
		// x[i]-lからx[i]-rの範囲にいるxを抽出し、その添字を知りたい
	}
}
