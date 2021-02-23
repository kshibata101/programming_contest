package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var cnt int

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	shellSort(a, n)
	printArrayBreak(a)
}

func shellSort(a []int, n int) {
	G := calcG(n)
	m := len(G)
	fmt.Println(m)
	printArray(G)
	for i := 0; i < m; i++ {
		insertionSort(a, n, G[i])
	}
	fmt.Println(cnt)
}

func calcG(n int) []int {
	g := []int{}
	v := 1
	for v <= n {
		g = append(g, v)
		v = 3*v + 1
	}
	m := len(g)
	G := make([]int, m)
	for i := 0; i < m; i++ {
		G[i] = g[m-i-1]
	}
	return G
}

func insertionSort(a []int, n int, g int) {
	for i := g; i < n; i++ {
		v := a[i]
		j := i - g
		for j >= 0 && a[j] > v {
			a[j+g] = a[j]
			j -= g
			cnt++
		}
		a[j+g] = v
	}
}

func printArray(a []int) {
	fmt.Print(a[0])
	for i := 1; i < len(a); i++ {
		fmt.Printf(" %d", a[i])
	}
	fmt.Println()
}

func printArrayBreak(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
