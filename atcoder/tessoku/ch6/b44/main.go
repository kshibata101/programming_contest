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

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	a := make([][]int, n)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = Scan()
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	q := Scan()
	for i := 0; i < q; i++ {
		c, x, y := Scan(), Scan()-1, Scan()-1
		if c == 1 {
			idx[x], idx[y] = idx[y], idx[x]
		} else {
			fmt.Fprintln(wr, a[idx[x]][y])
		}
	}
}
