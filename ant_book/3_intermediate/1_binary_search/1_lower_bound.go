package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	l := -1
	u := n
	for l+1 < u {
		m := (l+u)/2
		if a[m] >= k {
			u = m
		} else {
			l = m
		}
	}
	fmt.Println(u)
}