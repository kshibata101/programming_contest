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
	t := Scan()
	n := Scan()
	in := make([]int, t+1)
	out := make([]int, t+1)
	for i := 0; i < n; i++ {
		l := Scan()
		r := Scan()
		in[l] += 1
		out[r] += 1
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	num := 0
	for i := 0; i < t; i++ {
		num += in[i]
		num -= out[i]
		fmt.Fprintln(wr, num)
	}
}
