package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
	}
	return a
}

func ScanS() string {
	s.Scan()
	return s.Text()
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	n := Scan()
	r := Scan()
	p := 1000000007

	inv := make([]int, r+1)
	factInv := make([]int, r+1)
	inv[1] = 1
	factInv[0] = 1
	factInv[1] = 1
	for i := 2; i <= r; i++ {
		inv[i] = p - inv[p%i]*(p/i)%p
		factInv[i] = factInv[i-1] * inv[i] % p
	}
	ans := 1
	for i := n; i >= n-r+1; i-- {
		ans = (ans * i) % p
	}
	fmt.Println(ans * factInv[r] % p)
}
