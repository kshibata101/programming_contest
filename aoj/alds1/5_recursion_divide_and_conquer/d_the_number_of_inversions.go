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
	c := mergeSort(&a, 0, len(a))
	fmt.Println(c)
}


func mergeSort(a *[]int, l, r int) int {
	c := 0
	if l + 1 < r {
		m := (l + r) / 2
		c += mergeSort(a, l, m)
		c += mergeSort(a, m, r)
		c += merge(a, l, m, r)
	}
	return c
}

func merge(a *[]int, l, m, r int) int {
	// 番兵
	ll := append([]int{}, (*a)[l:m]...)
	ll = append(ll, 1 << 30)
	rr := append([]int{}, (*a)[m:r]...)
	rr = append(rr, 1 << 30)

	i, j, c := 0, 0, 0
	for k := l; k < r; k++ {
		if ll[i] <= rr[j] {
			(*a)[k] = ll[i]
			i++
		} else {
			(*a)[k] = rr[j]
			j++
			c += (m - l) - i
		}
	}
	return c
}
