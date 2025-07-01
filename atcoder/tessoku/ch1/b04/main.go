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

func ScanS() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := ScanS()
	ans := 0
	mul := 1
	for i := 0; i < len(n); i++ {
		if n[len(n)-i-1] == '1' {
			ans += mul
		}
		mul *= 2
	}
	fmt.Println(ans)
}
