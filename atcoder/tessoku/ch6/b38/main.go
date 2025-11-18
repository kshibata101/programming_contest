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

func ScanS() string {
	sc.Scan()
	return sc.Text()
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
	s := ScanS()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = 1
	}
	for i := 0; i < n-1; i++ {
		if s[i] == 'A' {
			h[i+1] = Max(h[i+1], h[i]+1)
		}
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == 'B' {
			h[i] = Max(h[i], h[i+1]+1)
		}
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += h[i]
	}
	fmt.Println(sum)
}
