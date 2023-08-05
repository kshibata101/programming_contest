package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i], _ = strconv.Atoi(sc.Text())
	}

	c := 0
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())

		for j := 0; j < n; j++ {
			if t == s[j] {
				c++
				break
			}
		}
	}
	fmt.Println(c)
}
