package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	c := 0
	for i := 0; i < n; i++ {
		var z int
		fmt.Scan(&z)

		if isPrime(z) {
			c++
		}
	}

	fmt.Println(c)
}

func isPrime(z int) bool {
	if z == 2 {
		return true
	}
	if z % 2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(z))); i = i+2 {
		if z % i == 0 {
			return (z / i < 2)
		}
	}
	return true
}
