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
	set := make(map[int]int)
	n := Scan()
	for i := 0; i < n; i++ {
		a := Scan()
		set[a]++
	}
	ans := 0
	for _, c := range set {
		if c > 1 {
			ans += c * (c - 1) / 2
		}
	}
	fmt.Println(ans)
}
