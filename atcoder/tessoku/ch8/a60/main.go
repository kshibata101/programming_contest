package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

type StockPrice struct {
	d int
	p int
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	n := Scan()
	sps := []StockPrice{}

	for i := 1; i <= n; i++ {
		a := Scan()
		cut := -1
		var kisan *StockPrice
		for j := len(sps) - 1; j >= 0; j-- {
			if a < sps[j].p {
				kisan = &sps[j]
				break
			}
			cut = j
		}
		if kisan == nil {
			fmt.Fprint(wr, -1)
		} else {
			fmt.Fprint(wr, kisan.d)
		}

		if cut >= 0 {
			sps = sps[:cut]
		}
		if i < n {
			fmt.Fprint(wr, " ")
		}
		sps = append(sps, StockPrice{i, a})
	}
	fmt.Fprintln(wr)
}
