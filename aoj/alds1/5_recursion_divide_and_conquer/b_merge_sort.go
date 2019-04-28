package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	s := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i], _ = strconv.Atoi(sc.Text())
	}

	cnt := mergeSort(&s, 0, len(s), 0)
	fmt.Print(s[0])
	for i := 1; i < n; i++ {
		fmt.Print(" ")
		fmt.Print(s[i])
	}
	fmt.Println()
	fmt.Println(cnt)
}

func mergeSort(a *[]int, l int, r int, cnt int) int {
	if l + 1 < r {
		m := (l + r) / 2
		cnt = mergeSort(a, l, m, cnt)
		cnt = mergeSort(a, m, r, cnt)
		cnt = merge(a, l, m, r, cnt)
	}
	return cnt
}

func merge(a *[]int, l int, m int, r int, cnt int) int {
	ll := append([]int{}, (*a)[l:m]...)
	ll = append(ll, 1 << 30)
	rr := append([]int{}, (*a)[m:r]...)
	rr = append(rr, 1 << 30)

	i := 0
	j := 0
	for k := l; k < r; k++ {
		if ll[i] <= rr[j] {
			(*a)[k] = ll[i]
			i++
		} else {
			(*a)[k] = rr[j]
			j++
		}
		cnt++
	}
	return cnt
}
