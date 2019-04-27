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
	sortWithPrint(n, a)
}

func sortWithPrint(n int, a []int) {
	for i := 0; i < n; i++ {
		v := a[i]
		j := i-1
		for ; j >= 0 && a[j] > v; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = v

		// print
		fmt.Print(a[0])
		for k := 1; k < n; k++ {
			fmt.Printf(" %d", a[k])
		}
		fmt.Println()
	}
}
