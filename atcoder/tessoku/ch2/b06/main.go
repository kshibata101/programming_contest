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
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a := Scan()
		sum[i] = sum[i-1] + a
	}
	q := Scan()
	for i := 0; i < q; i++ {
		l := Scan()
		r := Scan()
		atari := sum[r] - sum[l-1]
		if atari*2 < r-l+1 {
			fmt.Println("lose")
		} else if atari*2 > r-l+1 {
			fmt.Println("win")
		} else {
			fmt.Println("draw")
		}
	}
}
