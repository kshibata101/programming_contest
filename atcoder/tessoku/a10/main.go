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

	n := load(s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = load(s)
	}
	ml := make([]int, n)
	mr := make([]int, n)
	ml[0] = a[0]
	mr[n-1] = a[n-1]
	for i := 1; i < n; i++ {
		if ml[i-1] > a[i] {
			ml[i] = ml[i-1]
		} else {
			ml[i] = a[i]
		}

		if mr[n-i] > a[n-i-1] {
			mr[n-i-1] = mr[n-i]
		} else {
			mr[n-i-1] = a[n-i-1]
		}
	}

	d := load(s)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < d; i++ {
		l := load(s)
		r := load(s)

		max := ml[l-2]
		if mr[r] > max {
			max = mr[r]
		}
		fmt.Fprintln(w, max)
	}
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
