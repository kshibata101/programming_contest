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

	a, cnt := selectionSort(n, a)
	printArray(a)
	fmt.Println(cnt)
}

func selectionSort(n int, a []int) ([]int, int) {
	cnt := 0
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if a[j] < a[minj] {
				minj = j
			}
		}
		a[i], a[minj] = a[minj], a[i]
		if i != minj {
			cnt++
		}
	}
	return a, cnt
}

func printArray(a []int) {
	fmt.Print(a[0])
	for i := 1; i < len(a); i++ {
		fmt.Printf(" %d", a[i])
	}
	fmt.Println()
}
