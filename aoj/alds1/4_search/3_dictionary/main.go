package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	dic := make(map[string]bool)
	for i := 0; i < n; i++ {
		s.Scan()
		command := s.Text()
		s.Scan()
		word := s.Text()
		if command == "insert" {
			dic[word] = true
		} else if command == "find" {
			if _, ok := dic[word]; ok {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}
