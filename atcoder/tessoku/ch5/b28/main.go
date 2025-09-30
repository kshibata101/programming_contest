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
	a[0], a[1] = 1, 1
	for i := 2; i < n; i++ {
		a[i] = (a[i-1] + a[i-2]) % 1000000007
	}
	fmt.Println(a[n-1])
}
