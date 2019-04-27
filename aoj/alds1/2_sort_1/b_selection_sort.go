package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := []int{}
	for i := 0; i < n; i++ {
		var aa int
		fmt.Scan(&aa)

		a = append(a, aa)
	}

	c := 0
	for i := 0; i < n; i++ {
		minj := i
		for j := i+1; j < n; j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		if i != minj {
			a[i], a[minj] = a[minj], a[i]
			c++
		}
	}

	fmt.Print(a[0])
	for i := 1; i < n; i++ {
		fmt.Print(" ")
		fmt.Print(a[i])
	}
	fmt.Println()
	fmt.Println(c)
}
