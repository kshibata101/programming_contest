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

func ScanS() string {
	s.Scan()
	return s.Text()
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	p := map[string]int{}
	q := Scan()
	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			x := ScanS()
			y := Scan()
			p[x] = y
		} else if c == 2 {
			x := ScanS()
			fmt.Fprintln(w, p[x])
		} else {
			panic("")
		}
	}
}
