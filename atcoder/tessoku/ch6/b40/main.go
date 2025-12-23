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

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	cnt := make([]int, 100)
	for i := 0; i < n; i++ {
		a := Scan()
		cnt[a%100]++
	}
	sum := 0
	for i := 1; i < 50; i++ {
		sum += cnt[i] * cnt[100-i]
	}
	sum += cnt[0] * (cnt[0] - 1) / 2
	sum += cnt[50] * (cnt[50] - 1) / 2
	fmt.Println(sum)
}
