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
	n, k := Scan(), Scan()
	a := make([]int, k)
	for i := 0; i < k; i++ {
		a[i] = Scan()
	}
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for j := 0; j < k; j++ {
			if i >= a[j] && !dp[i-a[j]] {
				dp[i] = true
			}
		}
	}
	if dp[n] {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
