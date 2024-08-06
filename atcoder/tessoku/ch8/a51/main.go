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

func ScanS() string {
	s.Scan()
	return s.Text()
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	q := Scan()
	books := make([]string, 0, 100000)
	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			book := ScanS()
			books = append(books, book)
		} else if c == 2 {
			fmt.Println(books[len(books)-1])
		} else if c == 3 {
			books = books[:len(books)-1]
		} else {
			panic("invalid query")
		}
	}
}
