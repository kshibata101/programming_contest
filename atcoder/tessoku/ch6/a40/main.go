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

func Permutation(n, k int) int {
	ans := 1
	for i := 0; i < k; i++ {
		ans *= n - i
	}
	return ans
}

func Combination(n, k int) int {
	if n < k {
		return 0
	}
	if k == 0 {
		return 1
	}
	return Permutation(n, k) / Permutation(k, k)
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := Scan()
	a := map[int]int{}
	for i := 0; i < n; i++ {
		s := Scan()
		if _, ok := a[s]; !ok {
			a[s] = 0
		}
		a[s]++
	}
	ans := 0
	for _, v := range a {
		ans += Combination(v, 3)
	}
	fmt.Println(ans)
}
