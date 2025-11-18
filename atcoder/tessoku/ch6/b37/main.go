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

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var n uint64 = uint64(Scan())
	var sum uint64 = 0
	var p uint64 = 1
	for i := 0; i < 15; i++ { // iは桁(=10^iの位)
		var k uint64
		for k = 1; k <= 9; k++ {
			d := (n % (p * 10)) / p
			if k < d {
				sum += (n/(p*10) + 1) * p * k
			} else if k > d {
				sum += (n / (p * 10)) * p * k
			} else {
				sum += ((n/(p*10))*p + n%p + 1) * k
			}
		}
		p *= 10
	}
	fmt.Println(sum)
}
