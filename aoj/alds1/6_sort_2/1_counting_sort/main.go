package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	a = countingSort(a)
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}

func countingSort(a []int) []int {
	k := 10000
	C := make([]int, k+1)
	n := len(a)
	for i := 0; i < n; i++ {
		C[a[i]]++
	}
	for i := 0; i < k; i++ {
		C[i+1] += C[i]
	}

	b := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		b[C[a[i]]-1] = a[i]
		C[a[i]]--
	}
	return b
}
