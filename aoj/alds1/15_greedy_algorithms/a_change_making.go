package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	coins := []int{25,10,5,1}
	num := 0
	for i := 0; i < len(coins); i++ {
		cn := n / coins[i]
		n -= cn * coins[i]
		num += cn
	}
	fmt.Println(num)
}