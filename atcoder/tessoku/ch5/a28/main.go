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
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	m := 0
	n := ScanI()
	for i := 0; i < n; i++ {
		c := Scan()
		a := ScanI()
		if c == "+" {
			m += a
		} else if c == "-" {
			m -= a
		} else {
			m *= a
		}
		if m < 0 {
			m += 10000
		}
		m = m % 10000
		fmt.Fprintln(w, m)
	}
}
