package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	sc := bufio.NewScanner(os.Stdin)

	max, max2 := -1, -1
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
		if a[i] >= max {
			max2 = max
			max = a[i]
		} else if a[i] >= max2 {
			max2 = a[i]
		}
	}
	for i := 0; i < n; i++ {
		if a[i] == max {
			fmt.Println(max2)
		} else {
			fmt.Println(max)
		}
	}
}