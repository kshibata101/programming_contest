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
	n, m := Scan(), Scan()
	ans := make([]int, n)
	for i := 0; i < m; i++ {
		ans[Scan()-1]--
	}
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintln(wr, ans[i]+m)
	}
}
