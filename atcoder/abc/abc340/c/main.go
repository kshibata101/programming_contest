package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := load(s)
	a := make(map[int]int, n+1)
	fmt.Println(calc(n, a))
}

func calc(v int, a map[int]int) int {
	if v == 1 {
		return 0
	}
	_, ok := a[v];
	if ok && a[v] > 0 {
		return a[v]
	}
	res := v
	if v % 2 == 0 {
		res += calc(v/2, a) * 2
	} else {
		res += calc(v/2, a) + calc(v/2+1, a)
	}
	a[v] = res
	return res
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}