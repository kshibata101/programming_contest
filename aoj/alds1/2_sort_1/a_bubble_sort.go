package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scan(&n)

	arr := []int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	c := 0
	flag := 1
	for flag > 0 {
		flag = 0
		for j := n-1; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				c++
				flag = 1
			}
		}
	}
	fmt.Print(arr[0])
	for i := 1; i < n; i++ {
		fmt.Print(" ")
		fmt.Print(arr[i])
	}
	fmt.Println()
	fmt.Println(c)
}
