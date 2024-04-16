package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 9; i >= 0; i-- {
		if (n >> i) & 1 > 0 {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
	fmt.Println()
}