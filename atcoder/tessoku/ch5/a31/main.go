package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	ans += n / 3
	ans += n / 5
	ans -= n / 15
	fmt.Println(ans)
}
