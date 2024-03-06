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
	a := load(s)
	b := load(s)
	d := load(s)
	fmt.Print(a)
	for i := a + d; i <= b; i += d {
		fmt.Printf(" %d", i)
	}
	fmt.Println()
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}