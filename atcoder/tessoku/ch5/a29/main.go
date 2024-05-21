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

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	a := Scan()
	b := Scan()
	ans := 1
	m := 1000000007
	for i := 0; i < 30; i++ {
		beki := (1 << i)
		if (b/beki)%2 == 1 {
			ans = (ans * a) % m
		}
		a = (a * a) % m
	}
	fmt.Println(ans)
}
