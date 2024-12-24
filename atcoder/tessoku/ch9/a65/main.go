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

func Search(buka *[][]int, ans *[]int, i int) int {
	num := 0
	for j := 0; j < len((*buka)[i]); j++ {
		num += Search(buka, ans, (*buka)[i][j]) + 1
	}
	(*ans)[i] = num
	return num
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	buka := make([][]int, n+1)
	for i := 0; i < n; i++ {
		buka[i+1] = []int{}
	}
	for i := 2; i <= n; i++ {
		a := Scan()
		buka[a] = append(buka[a], i)
	}
	ans := make([]int, n+1)
	_ = Search(&buka, &ans, 1)

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	fmt.Fprint(wr, ans[1])
	for i := 2; i <= n; i++ {
		fmt.Fprintf(wr, " %d", ans[i])
	}
	fmt.Fprintln(wr)
}
