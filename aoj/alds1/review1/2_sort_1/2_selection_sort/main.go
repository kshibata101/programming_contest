package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	selectionSort(a, n)
}

func selectionSort(a []int, n int) {
	c := 0
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		a[i], a[minj] = a[minj], a[i]
		if i != minj {
			c++
		}
}
	fmt.Print(a[0])
	for i := 1; i < n; i++ {
		fmt.Printf(" %d", a[i])
	}
	fmt.Println()
	fmt.Println(c)
}