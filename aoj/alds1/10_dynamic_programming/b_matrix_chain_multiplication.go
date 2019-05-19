package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		fmt.Scan(&a[i+1])
	}

	m := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		m[i] = make([]int, n+1)
	}

	for d := 2; d <= n; d++ {
		for s := 0; s+d <= n; s++ {
			min := 1 << 30
			for j := s+1; j < s+d; j++ {
				val := m[s][j] + m[j][s+d] + a[s]*a[j]*a[s+d]
				if val < min {
					min = val
				}
			}
			m[s][s+d] = min
		}
	}
	fmt.Println(m[0][n])
}
