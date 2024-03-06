package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	c := 0
	for i := 1; c < n; i++ {
		for j := 1; j <= i; j++ {
			for k := 1; k <= j; k++ {
				c++
				if c >= n {
					fmt.Println(repu(i) + repu(j) + repu(k))
					return
				}
			}
		}
	}
}

func repu(m int) int {
	r := 1
	for i := 0; i < m - 1; i++ {
		r = r * 10 + 1
	}
	return r
}