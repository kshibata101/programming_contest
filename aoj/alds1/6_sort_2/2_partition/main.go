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
	i := partition(a, 0, n-1)

	if i > 0 {
		fmt.Print(strings.Trim(fmt.Sprint(a[:i]), "[]"))
		fmt.Print(" ")
	}
	fmt.Printf("[%d]", a[i])
	if i+1 < n {
		fmt.Print(" ")
		fmt.Print(strings.Trim(fmt.Sprint(a[i+1:]), "[]"))
	}
	fmt.Println()
}

func partition(a []int, p, r int) int {
	i := p - 1
	for j := p; j < r; j++ {
		if a[j] <= a[r] {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
