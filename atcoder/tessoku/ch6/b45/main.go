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

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a, b, c := Scan(), Scan(), Scan()
	if a+b+c == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
