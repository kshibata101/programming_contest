package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
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
	q := partition(&a, 0, len(a) - 1)

	if q > 0 {
		fmt.Print(strings.Trim(fmt.Sprint(a[:q]), "[]"))
		fmt.Print(" ")
	}
	fmt.Print("[")
	fmt.Print(a[q])
	fmt.Print("]")
	if q < len(a) - 1 {
		fmt.Print(" ")
		fmt.Print(strings.Trim(fmt.Sprint(a[q+1:]), "[]"))
	}
	fmt.Println()
}

func partition(a *[]int, p, r int) int {
	x := (*a)[r]
	i := p
	for j := p; j < r; j++ {
		if (*a)[j] <= x {
			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
			i++
		}
	}
	(*a)[i], (*a)[r] = (*a)[r], (*a)[i]
	return i
}
