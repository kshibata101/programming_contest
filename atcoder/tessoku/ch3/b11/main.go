package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	q := Scan()
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < q; i++ {
		x := Scan()
		l := -1
		r := n
		for r-l > 1 {
			m := (l + r) / 2
			if a[m] < x {
				l = m
			} else {
				r = m
			}
		}
		fmt.Fprintln(wr, r)
	}
}
