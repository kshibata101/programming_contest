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
	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := Scan()
			if v != 0 {
				x[j] = v
				y[i] = v
			}
		}
	}
	invX := 0
	invY := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if x[i] > x[j] {
				invX++
			}
			if y[i] > y[j] {
				invY++
			}
		}
	}
	fmt.Println(invX + invY)
}
