package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Print(a * b)
	fmt.Print(" ")
	fmt.Println(2 * (a + b))
}
