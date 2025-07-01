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
	xy := [1501][1501]int{}
	for i := 0; i < n; i++ {
		x := Scan()
		y := Scan()
		xy[x][y] += 1
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	sum := [1501][1501]int{}
	for i := 0; i < 1500; i++ {
		for j := 0; j < 1500; j++ {
			sum[i+1][j+1] = sum[i][j+1] + sum[i+1][j] - sum[i][j] + xy[i+1][j+1]
		}
	}

	q := Scan()
	for i := 0; i < q; i++ {
		a := Scan()
		b := Scan()
		c := Scan()
		d := Scan()
		fmt.Fprintln(wr, sum[c][d]-sum[a-1][d]-sum[c][b-1]+sum[a-1][b-1])
	}
}
