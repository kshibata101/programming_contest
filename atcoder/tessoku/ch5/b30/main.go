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

func Combination(n, k int, p int) int {
	if n <= 1 {
		return 1
	}
	fact := make([]int, n+1)
	inv := make([]int, n+1)
	factInv := make([]int, n+1)
	fact[0] = 1
	inv[0] = 1
	factInv[0] = 1
	fact[1] = 1
	inv[1] = 1
	factInv[1] = 1
	for i := 2; i <= n; i++ {
		fact[i] = fact[i-1] * i % p
		inv[i] = p - inv[p%i]*(p/i)%p
		factInv[i] = factInv[i-1] * inv[i] % p
	}
	return fact[n] * (factInv[k] * factInv[n-k] % p) % p
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	h := Scan()
	w := Scan()
	fmt.Println(Combination(h+w-2, h-1, 1000000007))
}
