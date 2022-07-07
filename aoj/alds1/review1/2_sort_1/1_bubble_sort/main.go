package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	bubbleSort(a, n)
}

func bubbleSort(a []int, n int) {
	c := 0
	flag := true
	for flag {
		flag = false
		for j := n-1; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				c++
				flag = true
			}
		}
	}
	fmt.Print(a[0])
	for i := 1; i < n; i++ {
		fmt.Printf(" %d", a[i])
	}
	fmt.Println()
	fmt.Println(c)
}