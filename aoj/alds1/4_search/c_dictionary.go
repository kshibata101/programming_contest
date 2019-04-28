package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	dic := make(map[string]struct{})
	for i := 0; i < n; i++ {
		s.Scan()
		arr := strings.Split(s.Text(), " ")

		if arr[0] == "insert" {
			dic[arr[1]] = struct{}{}
		} else if arr[0] == "find" {
			if _, ok := dic[arr[1]]; ok {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}

