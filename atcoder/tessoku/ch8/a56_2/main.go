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

func ScanS() string {
	s.Scan()
	return s.Text()
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	buf := make([]byte, 10000)
	s.Buffer(buf, 200001)
	n := Scan()
	q := Scan()
	S := ScanS()
	var mod int64 = 2147483647

	// Sを数値列に変換
	t := make([]int64, n+1)
	for i := 0; i < n; i++ {
		t[i+1] = int64(S[i]-'a') + 1
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	p := make([]int64, n+1) // power
	p[0] = 1
	h := make([]int64, n+1) // hash
	h[0] = 0
	for i := 0; i < n; i++ {
		p[i+1] = p[i] * 100 % mod
		h[i+1] = (h[i]*100 + t[i+1]) % mod
	}

	for i := 0; i < q; i++ {
		a := Scan()
		b := Scan()
		c := Scan()
		d := Scan()

		h1 := h[b] - (h[a-1] * p[b-a+1] % mod)
		h2 := h[d] - (h[c-1] * p[d-c+1] % mod)
		if h1 < 0 {
			h1 += mod
		}
		if h2 < 0 {
			h2 += mod
		}
		if h1 == h2 {
			fmt.Fprintln(wr, "Yes")
		} else {
			fmt.Fprintln(wr, "No")
		}
	}
}
