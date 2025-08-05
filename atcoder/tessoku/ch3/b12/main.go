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
	n := float64(Scan())
	l := 0.0
	r := 50.0
	for l < r {
		x := (l + r) / 2
		diff := x*x*x + x - n
		if diff < -0.001 {
			l = x
		} else if diff > 0.001 {
			r = x
		} else {
			fmt.Println(x)
			return
		}
	}
}
