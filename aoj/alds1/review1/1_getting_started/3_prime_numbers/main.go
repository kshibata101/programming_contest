package main

import (
	"fmt"
	"math"
)

func main() {
	var n, c, v int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if isPrimeNumber(v) {
			c++
		}
	}
	fmt.Println(c)
}

func isPrimeNumber(v int) bool {
	if v == 1 {
		return false
	} else if v == 2 {
		return true
	} else if v % 2 == 0 {
		return false
	}
	sqrtv := math.Sqrt(float64(v))
	for i := 3; i <= int(sqrtv); i=i+2 {
		if v % i == 0 {
			return false
		}
	}
	return true
}