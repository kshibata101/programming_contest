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
	l := Scan()
	ans := 0
	for i := 0; i < n; i++ {
		a := Scan()
		b := ScanS()
		d := 0
		if b == "E" {
			d = l - a
		} else if b == "W" {
			d = a
		}
		if d > ans {
			ans = d
		}
	}
	fmt.Println(ans)
}
