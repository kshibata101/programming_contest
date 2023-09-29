package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		s := "-"
		for j := 1; j <= 9; j++ {
			if n%j != 0 {
				continue
			}
			if i%(n/j) == 0 {
				s = strconv.Itoa(j)
				break
			}
		}
		fmt.Print(s)
	}
}
