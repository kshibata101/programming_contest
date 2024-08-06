package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
	}
	return a
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := Scan()
	nim := 0
	for i := 0; i < n; i++ {
		a := Scan()
		nim = nim ^ a
	}
	if nim == 0 {
		fmt.Println("Second")
	} else {
		fmt.Println("First")
	}
}
