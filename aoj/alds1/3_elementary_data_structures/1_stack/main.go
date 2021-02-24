package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intstack []int

func (s *intstack) push(i int) {
	*s = append(*s, i)
}

func (s *intstack) pop() int {
	i := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return i
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	st := s.Text()
	a := strings.Split(st, " ")
	var stack intstack = make([]int, 2)
	for i := 0; i < len(a); i++ {
		if a[i] == "+" {
			stack.push(stack.pop() + stack.pop())
		} else if a[i] == "-" {
			stack.push(-stack.pop() + stack.pop())
		} else if a[i] == "*" {
			stack.push(stack.pop() * stack.pop())
		} else {
			v, _ := strconv.Atoi(a[i])
			stack.push(v)
		}
	}
	fmt.Println(stack.pop())
}
