package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

// (l, r]は求めたい半開区間、(a, b]はuに対応する半開区間、uは現在のセグメントツリーのセル番号
func Query(arr []int, l, r, a, b, u int) int {
	if r <= a || b <= l {
		return 0
	}
	if l <= a && b <= r {
		return arr[u]
	}
	m := (a + b) / 2
	return Query(arr, l, r, a, m, 2*u) + Query(arr, l, r, m, b, 2*u+1)
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	q := Scan()

	size := 1
	for size < n {
		size *= 2
	}
	a := make([]int, 2*size)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			p := Scan() + size - 1
			x := Scan() - a[p]
			for p > 0 {
				a[p] += x
				p = p / 2
			}
		} else if c == 2 {
			l := Scan()
			r := Scan()
			ans := Query(a, l, r, 1, size+1, 1)
			fmt.Fprintln(wr, ans)
		} else {
			panic("")
		}
	}
}
