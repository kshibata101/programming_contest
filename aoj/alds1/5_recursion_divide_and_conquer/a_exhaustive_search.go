package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}
	
	set := make(map[int]struct{})
	max := int(math.Pow(2, float64(n)))
	for i := 1; i <= max; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if (i >> uint(j)) & 1 != 0 {
				sum += a[j]
			}
		}
		set[sum] = struct{}{}
	}

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())

		if _, ok := set[m]; ok {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
