package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() string {
	s.Scan()
	return s.Text()
}

func ScanI() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanIArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ScanI()
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	q := ScanI()
	for i := 0; i < q; i++ {
		if isPrime(ScanI()) {
			fmt.Fprintln(w, "Yes")
		} else {
			fmt.Fprintln(w, "No")
		}
	}
}

func isPrime(v int) bool {
	if v == 2 {
		return true
	}
	if v%2 == 0 {
		return false
	}
	for i := 3; float64(i) <= math.Sqrt(float64(v)); i += 2 {
		if v%i == 0 {
			return false
		}
	}
	return true
}
