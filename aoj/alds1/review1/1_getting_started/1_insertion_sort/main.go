package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	insertionSort(n, a)
}

func insertionSort(n int, a []int) {
	printArray(a)
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v
		printArray(a)
	}
}

func printArray(a []int) {
	fmt.Print(a[0])
	for i := 1; i < len(a); i++ {
		fmt.Print(" ")
		fmt.Print(a[i])
	}
	fmt.Println()
}