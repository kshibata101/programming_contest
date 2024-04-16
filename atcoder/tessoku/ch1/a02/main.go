package main

import "fmt"

func main() {
	var n, x, a int
	fmt.Scan(&n)
	fmt.Scan(&x)
	res := false
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if x == a {
			res = true
		}
	}
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}