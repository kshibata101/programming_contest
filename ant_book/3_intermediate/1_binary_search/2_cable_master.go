package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

func main() {
	var n,k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	l := make([]float64, n)
	t := 0.0
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		l[i], _ = strconv.ParseFloat(sc.Text(), 64)
		t = math.Max(t, l[i])
	}

	b := 0.0
	for b + 0.01 < t {
		m := (b+t) / 2
		K := calc(m, &l)
		if K >= k {
			b = m
		} else {
			t = m
		}
	}
	fmt.Printf("%.2f\n", b)
}

func calc(s float64, l *[]float64) int {
	num := 0
	for i := 0; i < len(*l); i++ {
		num += int((*l)[i]/s)
	}
	return num
}