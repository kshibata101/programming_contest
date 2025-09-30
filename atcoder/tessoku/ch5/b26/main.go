package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func IsPrime(a int) bool {
	if a == 2 {
		return true
	}
	if a%2 == 0 {
		return false
	}
	v := int(math.Sqrt(float64(a)))
	for i := 3; i <= v; i += 2 {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			fmt.Fprintln(wr, i)
		}
	}
}
