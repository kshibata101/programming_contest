package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() string {
	s.Scan()
	return s.Text()
}

func ScanI() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanIArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ScanI()
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
	a := ScanI()
	b := ScanI()
	mx := Max(a, b)
	mn := Min(a, b)
	for mx > mn {
		if mx%mn == 0 {
			fmt.Println(mn)
			return
		}
		mx, mn = mn, mx%mn
	}
}
