package main

import "fmt"

func main() {
	var a, b uint64
	fmt.Scan(&a)
	fmt.Scan(&b)

	var c uint64 = 0
	for a != b {
		if a > b {
			c += a / b
			a = a % b
		} else {
			c += b / a
			b = b % a
		}
		if a == 0 || b == 0 {
			c -= 1
			break
		}
	}
	fmt.Println(c)
}
