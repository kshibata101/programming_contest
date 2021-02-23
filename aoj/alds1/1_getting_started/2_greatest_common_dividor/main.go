package main

import "fmt"

func main() {
	var x, y, l, s int
	fmt.Scan(&x)
	fmt.Scan(&y)

	if x < y {
		l, s = y, x
	} else {
		l, s = x, y
	}

	for l%s > 0 {
		l, s = s, l%s
	}
	fmt.Println(s)
}
