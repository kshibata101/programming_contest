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

	h := load(s)
	w := load(s)
	sum := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		sum[i] = make([]int, w+1)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			x := load(s)
			sum[i+1][j+1] = x + sum[i][j+1] + sum[i+1][j] - sum[i][j]
		}
	}

	q := load(s)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i := 0; i < q; i++ {
		a := load(s)
		b := load(s)
		c := load(s)
		d := load(s)
		fmt.Fprintln(wr, sum[c][d]-sum[a-1][d]-sum[c][b-1]+sum[a-1][b-1])
	}
}

func load(s *bufio.Scanner) int {
	s.Scan()
	res, _ := strconv.Atoi(s.Text())
	return res
}
