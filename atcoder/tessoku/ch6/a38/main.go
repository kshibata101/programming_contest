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
	d := Scan()
	n := Scan()
	m := make([]int, d)
	for i := 0; i < d; i++ {
		m[i] = 24
	}
	for i := 0; i < n; i++ {
		l := Scan() - 1
		r := Scan() - 1
		h := Scan()
		for j := l; j <= r; j++ {
			m[j] = Min(m[j], h)
		}
	}
	ans := 0
	for i := 0; i < d; i++ {
		ans += m[i]
	}
	fmt.Println(ans)
}
