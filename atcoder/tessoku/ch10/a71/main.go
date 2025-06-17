package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func ScanArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
	}
	return a
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	a := ScanArr(n)
	b := ScanArr(n)
	sort.Ints(a)
	sort.Ints(b)

	ans := 0
	for i := 0; i < n; i++ {
		ans += a[i] * b[n-i-1]
	}
	fmt.Println(ans)
}
