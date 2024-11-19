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

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// [l, r)は求めたい半開区間、uが現在のセグメント木のセル番号、[a, b)はuに対応する半開区間
func Query(arr []int, l, r, a, b, u int) int {
	if r <= a || b <= l {
		return -1
	}
	if l <= a && b <= r {
		return arr[u]
	}
	m := (a + b) / 2
	al := Query(arr, l, r, a, m, u*2)
	ar := Query(arr, l, r, m, b, u*2+1)
	return Max(al, ar)
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := Scan()
	size := 1
	for size < n {
		size *= 2
	}
	a := make([]int, size*2)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	q := Scan()
	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			pos := Scan()
			pos += size - 1
			x := Scan()
			a[pos] = x
			for pos > 1 {
				pos /= 2
				a[pos] = Max(a[pos*2], a[pos*2+1])
			}
		} else if c == 2 {
			l := Scan()
			r := Scan()
			ans := Query(a, l, r, 1, size+1, 1)
			fmt.Fprintln(w, ans)
		} else {
			panic("")
		}
	}
}
