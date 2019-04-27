package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := []string{}
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		arr = append(arr, str)
	}
	
	arr2 := make([]string, len(arr))
	copy(arr2, arr)

	// bubble sort
	flag := 1
	for flag > 0 {
		flag = 0
		for j := n-1; j > 0; j-- {
			if arr[j-1][1] > arr[j][1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
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
	fmt.Println("Stable")

	// selection sort
	for i := 0; i < n; i++ {
		minj := i
		for j := i+1; j < n; j++ {
			if arr2[j][1] < arr2[minj][1] {
				minj = j
			}
		}
		arr2[i], arr2[minj] = arr2[minj], arr2[i]
	}

	stable := true
	fmt.Print(arr2[0])
	stable = stable && arr[0] == arr2[0]
	for i := 1; i < n; i++ {
		fmt.Print(" ")
		fmt.Print(arr2[i])
		stable = stable && arr[i] == arr2[i]
	}
	fmt.Println()
	if stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}
