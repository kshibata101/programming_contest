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
		a := Scan()
		b := Scan()
		c := Scan()
		d := Scan()
		xy[a][b] += 1
		xy[a][d] -= 1
		xy[c][b] -= 1
		xy[c][d] += 1
	}
	sum := [1501][1501]int{}
	for i := 0; i <= 1500; i++ {
		for j := 0; j <= 1500; j++ {
			sum[i][j] = xy[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}
	ans := 0
	for i := 0; i <= 1500; i++ {
		for j := 0; j <= 1500; j++ {
			if sum[i][j] > 0 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
