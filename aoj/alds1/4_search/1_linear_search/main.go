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

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	S := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		S[i], _ = strconv.Atoi(s.Text())
	}

	s.Scan()
	q, _ := strconv.Atoi(s.Text())
	C := 0
	for i := 0; i < q; i++ {
		s.Scan()
		t, _ := strconv.Atoi(s.Text())

		for j := 0; j < n; j++ {
			if t == S[j] {
				C++
				break
			}
		}
	}
	fmt.Println(C)
}
