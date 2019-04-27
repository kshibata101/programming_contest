package main

import (
	"fmt"
)

func main() {
	var x, y int32
	fmt.Scan(&x)
	fmt.Scan(&y)

	if x < y {
		fmt.Println(gcd(y, x))
	} else {
		fmt.Println(gcd(x, y))
	}
	
}

func gcd(b, s int32) int32 {
	if b % s == 0 {
		return s
	}
	return gcd(s, b - s * (b / s))
}
