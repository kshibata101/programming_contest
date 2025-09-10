package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func Min(a, b float64) float64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	x := make([]int, n)
	y := make([]int, n)
	d := make([][]float64, n)
	for i := 0; i < n; i++ {
		x[i] = Scan()
		y[i] = Scan()
		d[i] = make([]float64, n)
		for j := 0; j < i; j++ {
			d[i][j] = math.Sqrt(float64((x[i]-x[j])*(x[i]-x[j]) + (y[i]-y[j])*(y[i]-y[j])))
			d[j][i] = d[i][j]
		}
	}

	// init dp
	size := 1 << n
	inf := float64(1 << 62)
	// 1st: a set of cities, 2nd: an index of current city
	dp := make([][]float64, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = inf
		}
	}
	// 都市0を起点とする
	dp[0][0] = 0

	// proceed
	for i := 0; i < size; i++ { // i: すでに訪れた都市の集合(j含まず)のbit表現
		for j := 0; j < n; j++ { // j: 最後にいる都市
			if i&(1<<j) > 0 { // jがiに含まれるならskip
				continue
			}
			nexti := i | (1 << j)
			for k := 0; k < n; k++ { // k: 次に訪れる都市
				if nexti&(1<<k) > 0 && nexti != size-1 { // すでにいったことあるならskip. ただし最後1週して戻るケースだけはok
					continue
				}
				dp[nexti][k] = Min(dp[nexti][k], dp[i][j]+d[j][k])
			}
		}
	}
	fmt.Println(dp[size-1][0])
}
