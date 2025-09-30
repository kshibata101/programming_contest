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
	a, b := Scan(), Scan()
	ans := 1
	mod := 1000000007
	for i := 0; i < 60; i++ {
		beki := (1 << i)
		if (b/beki)%2 == 1 {
			ans = (ans * a) % mod
		}
		a = (a * a) % mod
	}
	fmt.Println(ans)
}
