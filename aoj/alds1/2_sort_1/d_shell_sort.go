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

	a := []int{}
	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		s.Scan()
		aa, _ := strconv.Atoi(s.Text())
		a = append(a, aa)
	}

	shellSort(n, a)
}

func shellSort(n int, a []int) {
	cnt := 0

	g := []int{}
	for i := 2; i <= n+1; i *= 2 {
		g = append([]int{i-1}, g...)
	}
	m := len(g)

	for i := 0; i < m; i++ {
		var c int
		a, c = insertionSort(n, a, g[i])
		cnt += c
	}

	fmt.Println(m)
	fmt.Print(g[0])
	for i := 1; i < m; i++ {
		fmt.Print(" ")
		fmt.Print(g[i])
	}
	fmt.Println()
	fmt.Println(cnt)
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}

func insertionSort(n int, a []int, g int) ([]int, int) {
	cnt := 0
	for i := g; i < n; i++ {
		v := a[i]
		j := i-g
		for ; j >= 0; j = j-g {
			if a[j] > v {
				a[j+g] = a[j]
				cnt++
			} else {
				break
			}
		}
		a[j+g] = v
	}
	return a, cnt
}
