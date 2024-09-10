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
	q := Scan()
	que := make([]string, 0, 100000)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < q; i++ {
		c := Scan()
		if c == 1 {
			x := ScanS()
			que = append(que, x)
		} else if c == 2 {
			fmt.Fprintln(w, que[0])
		} else if c == 3 {
			que = que[1:]
		} else {
			panic("undefined command")
		}
	}
}
