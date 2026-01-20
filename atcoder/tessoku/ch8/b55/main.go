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

func BinSearch(a []int, x int) int {
	l, r := 0, len(a)
	for l < r {
		m := (l + r) / 2
		if x < a[m] {
			r = m
		} else if x > a[m] {
			l = m + 1
		} else {
			return m
		}
	}
	return l
}

func Dis(a, b int) int {
	return Max(a, b) - Min(a, b)
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	q := Scan()
	a := make([]int, 0, q)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < q; i++ {
		c, x := Scan(), Scan()
		if c == 1 {
			pos := BinSearch(a, x)
			a = append(a, 0)
			copy(a[pos+1:], a[pos:])
			a[pos] = x
		} else if c == 2 {
			if len(a) == 0 {
				fmt.Fprintln(wr, -1)
				continue
			}

			df := 1000000000
			pos := BinSearch(a, x)
			if pos < len(a) {
				df = Min(df, Dis(x, a[pos]))
			}
			if pos > 0 {
				df = Min(df, Dis(x, a[pos-1]))
			}
			fmt.Fprintln(wr, df)
		} else {
			panic("")
		}
	}
}
