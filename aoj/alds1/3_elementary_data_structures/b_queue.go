package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Process struct {
	Name string
	Time int
}

func main() {
	var n, q int
	fmt.Scan(&n)
	fmt.Scan(&q)

	queue := make([]Process, n)
	s := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		s.Scan()
		splits := strings.Split(s.Text(), " ")
		time, _ := strconv.Atoi(splits[1])
		queue[i] = Process{splits[0], time}
	}

	time := 0
	for ;len(queue) > 0; {
		process := queue[0]
		queue = queue[1:]

		if process.Time <= q {
			time += process.Time
			fmt.Print(process.Name)
			fmt.Print(" ")
			fmt.Println(time)
		} else {
			time += q
			process.Time -= q
			queue = append(queue, process)
		}
	}
}
