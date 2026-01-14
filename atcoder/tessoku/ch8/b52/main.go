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
	init := 10000
	max := 200000
	buf := make([]byte, init)
	sc.Buffer(buf, max)
	n, x := Scan(), Scan()
	a := ScanS()
	b := []byte(a)
	que := []int{x}
	b[x-1] = '@'
	for len(que) > 0 {
		pos := que[0]
		if len(que) > 1 {
			que = que[1:]
		} else {
			que = []int{}
		}
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
