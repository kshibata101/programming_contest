package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	if k >= n*2-2 && k%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
