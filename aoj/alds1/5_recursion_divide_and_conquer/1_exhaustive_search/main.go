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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		a[i], _ = strconv.Atoi(s.Text())
	}

	set := make(map[int]struct{})
	for j := 0; j < (1 << uint(n)); j++ {
		sum := 0
		for k := 0; k < n; k++ {
			if (j>>uint(k))&1 == 1 {
				sum += a[k]
			}
		}
		set[sum] = struct{}{}
	}

	s.Scan()
	q, _ := strconv.Atoi(s.Text())
	for i := 0; i < q; i++ {
		s.Scan()
		m, _ := strconv.Atoi(s.Text())

		if _, ok := set[m]; ok {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
