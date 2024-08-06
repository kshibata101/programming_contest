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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = i
	}
	q := Scan()
	normal := true
	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			x := Scan()
			y := Scan()
			if normal {
				a[x] = y
			} else {
				a[n+1-x] = y
			}
		} else if c == 2 {
			normal = !normal
		} else if c == 3 {
			x := Scan()
			if normal {
				fmt.Println(a[x])
			} else {
				fmt.Println(a[n+1-x])
			}
		}
	}
}
