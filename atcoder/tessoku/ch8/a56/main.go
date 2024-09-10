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
	buf := make([]byte, 10000)
	s.Buffer(buf, 200001)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	_ = Scan()
	q := Scan()
	S := ScanS()
	for i := 0; i < q; i++ {
		a := Scan()
		b := Scan()
		c := Scan()
		d := Scan()
		if S[a-1:b] == S[c-1:d] {
			fmt.Fprintln(w, "Yes")
		} else {
			fmt.Fprintln(w, "No")
		}
	}
}
