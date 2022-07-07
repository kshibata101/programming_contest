package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n, r int
	fmt.Scan(&n)
	fmt.Scan(&r)
	min := r
	diff := math.MinInt32
	s := bufio.NewScanner(os.Stdin)
	for i := 1; i < n; i++ {
		s.Scan()
		r, _ = strconv.Atoi(s.Text())
		if r - min > diff {
			diff = r - min
		}
		if r < min {
			min = r
		}
	}
	fmt.Println(diff)
}