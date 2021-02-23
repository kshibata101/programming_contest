package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n, mn, df int
	fmt.Scan(&n)
	fmt.Scan(&mn)
	df = math.MinInt32

	s := bufio.NewScanner(os.Stdin)
	for i := 1; i < n; i++ {
		s.Scan()
		r, _ := strconv.Atoi(s.Text())

		if r-mn > df {
			df = r - mn
		}
		if r < mn {
			mn = r
		}
	}
	fmt.Println(df)
}
