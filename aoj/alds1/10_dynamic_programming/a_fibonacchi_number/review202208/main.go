package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	a[0] = 1
	a[1] = 1
	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	fmt.Println(a[n])
}
