package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a := make([]int, n)
	k := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
		if a[i] > k {
			k = a[i]
		}
	}
	
	c := make([]int, k+1)
	for i := 0; i < n; i++ {
		c[a[i]]++
	}

	for i := 0; i < k; i++ {
		c[i+1] += c[i]
	}

	b := make([]int, n)
	for i := n-1; i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]]--
	}

	fmt.Println(strings.Trim(fmt.Sprint(b), "[]"))
}

