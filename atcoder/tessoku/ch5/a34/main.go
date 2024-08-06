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
	x := Scan()
	y := Scan()
	g := make([]int, 100001)
	for i := Min(x, y); i < 100001; i++ {
		gx, gy := -1, -1
		if i >= x {
			gx = g[i-x]
		}
		if i >= y {
			gy = g[i-y]
		}
		if g[i] == Min(gx, gy) {
			g[i]++
		}
		if g[i] == Max(gx, gy) {
			g[i]++
		}
	}
	xor := 0
	for i := 0; i < n; i++ {
		a := Scan()
		xor = xor ^ g[a]
	}
	if xor == 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
