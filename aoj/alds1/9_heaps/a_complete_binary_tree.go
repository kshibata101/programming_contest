package main

import (
	"fmt"
)

func main() {
	var h int
	fmt.Scan(&h)

	heap := make([]int, h+1)
	heap[0] = -1
	for i := 1; i <= h; i++ {
		fmt.Scan(&heap[i])
	}
	for i := 1; i <= h; i++ {
		fmt.Printf("node %d: key = %d, ", i, heap[i])
		if i > 1 {
			fmt.Printf("parent key = %d, ", heap[i/2])
		}
		if 2*i <= h {
			fmt.Printf("left key = %d, ", heap[2*i])
		}
		if 2*i+1 <= h {
			fmt.Printf("right key = %d, ", heap[2*i+1])
		}
		fmt.Println()
	}
}
