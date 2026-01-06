package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc *bufio.Scanner

func ScanS() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	buf := make([]byte, 10000)
	sc.Buffer(buf, 1000000)
	sc.Split(bufio.ScanWords)
	s := ScanS()
	stack := []int{}
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i+1)
		} else {
			l := stack[len(stack)-1] // last
			fmt.Fprintf(wr, "%d %d\n", l, i+1)
			stack = stack[:len(stack)-1] // pop
		}
	}
}
