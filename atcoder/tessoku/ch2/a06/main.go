package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, q, a int
	fmt.Scan(&n)
	fmt.Scan(&q)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())
		sum[i+1] = sum[i] + a
	}

	var l, r int
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < q; i++ {
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		fmt.Fprintln(w, sum[r]-sum[l-1])
	}
}
