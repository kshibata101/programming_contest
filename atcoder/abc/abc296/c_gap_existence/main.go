package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, x int
	fmt.Scan(&n)
	fmt.Scan(&x)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	a := make([]int, n)
	set := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
		set[a[i]] = struct{}{}
	}
	for i := 0; i < n; i++ {
		_, ok := set[a[i] + x]
		if ok {
			fmt.Println("Yes")
			return
		}
		_, ok = set[a[i] - x]
		if ok {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}