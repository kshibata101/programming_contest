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
	cap := 512 * 1024
	sc.Buffer(make([]byte, cap), cap)
	sc.Split(bufio.ScanWords)
	n, k := Scan(), Scan()
	s := ScanS()
	c := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			c++
		}
	}
	if (c+k)%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
