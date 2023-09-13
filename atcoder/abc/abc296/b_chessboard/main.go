package main

import "fmt"

func main() {
	var s string
	for i := 0; i < 8; i++ {
		fmt.Scan(&s)
		for j := 0; j < 8; j++ {
			if s[j] == '*' {
				fmt.Printf("%s%d\n", string('a' + j), 8 - i)
			}
		}
	}
}