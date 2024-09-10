package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func find(x int, arr []int) int {
	l := 0
	r := len(arr)
	for l < r {
		m := (l + r) / 2
		if x < arr[m] {
			r = m
		} else if x > arr[m] {
			l = m + 1
		} else {
			return m
		}
	}
	return l
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	q := Scan()
	arr := []int{}
	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			x := Scan()
			idx := find(x, arr)
			if idx >= len(arr) {
				arr = append(arr, x)
			} else {
				arr = append(arr[:idx], append([]int{x}, arr[idx:]...)...)
			}
		} else if c == 2 {
			x := Scan()
			idx := find(x, arr)
			if idx >= len(arr) {
				arr = arr[:idx-1]
			} else {
				arr = append(arr[:idx], arr[idx+1:]...)
			}
		} else if c == 3 {
			x := Scan()
			idx := find(x, arr)
			if idx >= len(arr) {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, arr[idx])
			}
		} else {
			panic("")
		}
	}
}
