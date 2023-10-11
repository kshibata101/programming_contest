package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	li := make([]int, n)
	sumL := 0
	for i := 0; i < n; i++ {
		s.Scan()
		li[i], _ = strconv.Atoi(s.Text())
		sumL += li[i]
	}

	l := 1
	r := sumL + n - 1
	for l < r {
		w := (l + r) / 2
		if check(w, m, li) {
			r = w
		} else {
			l = w + 1
		}
	}
	fmt.Println(r)
}

func check(w, m int, li []int) bool {
	n := len(li)
	r := 0
	l := 0
	for i := 0; i < n; i++ {
		if li[i] > w { // 単語が幅を超えたら確実にアウト
			return false
		}
		if li[i]+1 <= l {
			l = l - li[i] - 1
		} else {
			r++
			if r > m {
				return false
			}
			l = w - li[i]
		}

	}
	return true
}
