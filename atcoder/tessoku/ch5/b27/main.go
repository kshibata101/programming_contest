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

func Gcd(a, b int) int {
	l, s := a, b
	if a < b {
		l, s = b, a
	}
	for s > 1 {
		v := l % s
		if v == 0 {
			return s
		} else {
			l, s = s, v
		}
	}
	return 1
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a, b := Scan(), Scan()
	gcd := Gcd(a, b)
	fmt.Println(a / gcd * b)
}
