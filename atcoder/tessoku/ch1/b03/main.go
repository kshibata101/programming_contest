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
	n := Scan()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = Scan()
	}
	result := "No"
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if a[i]+a[j]+a[k] == 1000 {
					result = "Yes"
					break
				}
			}
		}
	}
	fmt.Println(result)
}
