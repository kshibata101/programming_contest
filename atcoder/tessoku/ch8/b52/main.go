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

func ScanS() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n, x := Scan(), Scan()
	s := ScanS()
	b := []byte(s)
	que := []int{x}
	for len(que) > 0 {
		l := len(que)
		pos := que[l-1]
		que = que[:l-1]
		if pos > 1 && b[pos-2] == '.' {
			b[pos-2] = '@'
			que = append(que, pos-1)
		}
		if pos < n && b[pos] == '.' {
			b[pos] = '@'
			que = append(que, pos+1)
		}
	}
	fmt.Println(string(b))
}
