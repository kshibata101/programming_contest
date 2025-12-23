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
	x, y := Scan(), Scan()
	his := [][]int{}
	for x+y > 2 {
		if x > y {
			his = append(his, []int{x, y})
			x -= y
		} else {
			his = append(his, []int{x, y})
			y -= x
		}
	}
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	n := len(his)
	fmt.Fprintln(wr, n)
	for i := n - 1; i >= 0; i-- {
		fmt.Fprintf(wr, "%d %d\n", his[i][0], his[i][1])
	}
}
