package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var d, n, l, r int
	fmt.Scan(&d)
	fmt.Scan(&n)
	diff := make([]int, d+1)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		l, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		diff[l-1] += 1
		diff[r] -= 1
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	sum := 0
	for i := 0; i < d; i++ {
		sum += diff[i]
		fmt.Fprintln(w, sum)
	}
}
