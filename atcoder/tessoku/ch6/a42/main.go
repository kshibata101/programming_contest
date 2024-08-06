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

func ScanS() string {
	s.Scan()
	return s.Text()
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
	k := Scan()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
		b[i] = Scan()
	}
	ans := 0
	for i := 1; i+k <= 100; i++ {
		for j := 1; j+k <= 100; j++ {
			cnt := 0
			for l := 0; l < n; l++ {
				if i <= a[l] && a[l] <= i+k && j <= b[l] && b[l] <= j+k {
					cnt++
				}
			}
			if cnt > ans {
				ans = cnt
			}
		}
	}
	fmt.Println(ans)
}
