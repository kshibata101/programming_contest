package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	num := 0

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		if isPrimeNumber(a) {
			num++
		}
	}
	fmt.Println(num)
}

func isPrimeNumber(x int) bool {
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(x))); i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
