package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := []int{}
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		a = append(a, tmp)
	}
	insertionSort(n, a)
}

func insertionSort(n int, a []int) {
	printArray(a)
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0 && a[j] > v; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = v
		printArray(a)
	}
}

func printArray(a []int) {
	fmt.Print(a[0])
	for i := 1; i < len(a); i++ {
		fmt.Printf(" %d", a[i])
	}
	fmt.Println()
}
