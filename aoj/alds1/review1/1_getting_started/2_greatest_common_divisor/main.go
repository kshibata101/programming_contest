package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	var res int
	if x > y {
		res = gcd(x, y)
	} else {
		res = gcd(y, x)
	}
	fmt.Println(res)
}

func gcd(l, s int) int {
	if l % s == 0 {
		return s
	}
	return gcd(s, l % s)
}