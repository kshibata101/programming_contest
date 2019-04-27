package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	list := []int{}
	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		s.Scan()
		splits := strings.Split(s.Text(), " ")
		command := splits[0]
		var key int
		if len(splits) == 2 {
			key, _ = strconv.Atoi(splits[1])
		}

		switch command {
		case "insert":
			list = append([]int{key}, list...)
		case "deleteFirst":
			list = list[1:]
		case "deleteLast":
			list = list[0:len(list)-1]
		case "delete":
			for i := 0; i < len(list); i++ {
				if list[i] == key {
					list = append(list[:i], list[i+1:]...)
					break
				}
			}
		}
	}

	fmt.Print(list[0])
	for i := 1; i < len(list); i++ {
		fmt.Print(" ")
		fmt.Print(list[i])
	}
	fmt.Println()
}
