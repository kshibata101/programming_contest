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
	a := Scan()
	b := Scan()
	result := "No"
	for i := a; i <= b; i++ {
		if 100%i == 0 {
			result = "Yes"
			break
		}
	}
	fmt.Println(result)
}
