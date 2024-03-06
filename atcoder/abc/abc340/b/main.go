package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	n := load(s)
	a := []int{}
	for i := 0; i < n; i++ {
		t := load(s)
		x := load(s)
		if t == 1 {
			a = append(a, x)
		} else {
			fmt.Println(a[len(a) - x])
		}
	}
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}