package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cnt int

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}
	mergeSort(a, 0, n)

	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
	fmt.Println(cnt)
}

func merge(a []int, l, m, r int) {
	L := make([]int, m-l)
	R := make([]int, r-m)
	copy(L, a[l:m])
	copy(R, a[m:r])
	L = append(L, 1<<31)
	R = append(R, 1<<31)

	i, j := 0, 0
	for k := l; k < r; k++ {
		if L[i] <= R[j] {
			a[k] = L[i]
			i++
		} else {
			a[k] = R[j]
			j++
		}
		cnt++
	}
}

func mergeSort(a []int, l, r int) {
	if l+1 >= r {
		return
	}
	m := (l + r) / 2
	mergeSort(a, l, m)
	mergeSort(a, m, r)
	merge(a, l, m, r)
}
