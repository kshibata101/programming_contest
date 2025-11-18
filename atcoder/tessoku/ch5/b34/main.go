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
		a := Scan()
		g := 2
		if a%5 <= 1 {
			g = 0
		} else if a%5 <= 3 {
			g = 1
		}
		nim = nim ^ g
	}
	if nim == 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
