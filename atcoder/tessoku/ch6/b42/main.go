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
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
		b[i] = Scan()
	}
	mul := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	ans := 0
	for i := 0; i < len(mul); i++ {
		sum := 0
		for j := 0; j < n; j++ {
			row := a[j]*mul[i][0] + b[j]*mul[i][1]
			if row > 0 {
				sum += row
			}
		}
		if sum > ans {
			ans = sum
		}
	}
	fmt.Println(ans)
}
