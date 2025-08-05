package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func makeSumArr(a []int) []int {
	n := len(a)
	sum := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		bit := i
		for j := 0; j < n; j++ {
			if bit&1 == 1 {
				sum[i] += a[j]
			}
			bit /= 2
		}
	}
	return sum
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	k := Scan()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
	}
	mae := makeSumArr(a[:n/2])
	ato := makeSumArr(a[n/2:])
	sort.Slice(mae, func(i, j int) bool { return mae[i] < mae[j] })
	sort.Slice(ato, func(i, j int) bool { return ato[i] < ato[j] })
	result := false
	for i := 0; i < 1<<(n/2); i++ {
		v := k - mae[i]
		if v == 0 {
			result = true
			break
		} else if v < 0 {
			continue
		}
		l := 0
		r := 1 << (n - n/2)
		for l < r {
			m := (l + r) / 2
			if ato[m] < v {
				l = m + 1
			} else if ato[m] > v {
				r = m
			} else {
				result = true
				break
			}
		}
		if result {
			break
		}
	}
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
