package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	arr := strings.Split(s.Text(), " ")
	stack := []int{}
	for _, s := range arr {
		if s == "+" || s == "-" || s == "*" {
			n1 := stack[len(stack)-2]
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch s {
			case "+":
				stack = append(stack, n1 + n2)
			case "-":
				stack = append(stack, n1 - n2)
			case "*":
				stack = append(stack, n1 * n2)
			}
		} else {
			n, _ := strconv.Atoi(s)
			stack = append(stack, n)
		}
	}
	fmt.Println(stack[0])
}
