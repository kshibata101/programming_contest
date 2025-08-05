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

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	k := Scan()
	a := make([]int, n+1)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = Scan()
		sum[i] = sum[i-1] + a[i]
	}

	ans := 0
	done := 0
	for i := 1; i <= n; i++ {
		l := i - 1
		r := n + 1
		for r-l > 1 {
			m := (l + r) / 2
			wa := sum[m] - sum[i-1]
			if wa < k {
				l = m
			} else if wa > k {
				r = m
			} else {
				l = m
				r = m + 1
			}
		}
		ans += (l - i + 1) * (l - i + 2) / 2
		if i <= done {
			ans -= (done - i + 1) * (done - i + 2) / 2
		}
		done = l
	}
	fmt.Println(ans)
}
