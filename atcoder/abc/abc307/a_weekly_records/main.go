package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	sum := 0
	for i := 0; i < 7*n; i++ {
		s.Scan()
		a, _ := strconv.Atoi(s.Text())
		sum += a
		if (i+1)%7 == 0 {
			if i > 6 {
				fmt.Print(" ")
			}
			fmt.Print(sum)
			sum = 0
		}
	}
}
