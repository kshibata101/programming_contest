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
	n, _, _ := Scan(), Scan(), Scan()
	nim := 0
	for i := 0; i < n; i++ {
		a, b := Scan()-1, Scan()-1
		nim = nim ^ a
		nim = nim ^ b
	}
	if nim > 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
