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
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	a := []int{}
	for i := 0; i < n; i++ {
		s.Scan()
		tmp, _ := strconv.Atoi(s.Text())
		a = append(a, tmp)
	}

	b, cnt := bubbleSort(n, a)
	printArray(b)
	fmt.Println(cnt)
}

func bubbleSort(n int, a []int) ([]int, int) {
	flag := true
	cnt := 0
	for flag {
		flag = false
		for j := n - 1; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
				flag = true
				cnt++
			}
		}
	}
	return a, cnt
}

func printArray(a []int) {
	fmt.Print(a[0])
	for i := 1; i < len(a); i++ {
		fmt.Printf(" %d", a[i])
	}
	fmt.Println()
}
