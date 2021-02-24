package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type process struct {
	name string
	time int
}

type processQueue []*process

func (q *processQueue) push(p *process) {
	*q = append(*q, p)
}

func (q *processQueue) pop() *process {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func main() {
	var n, q int
	fmt.Scan(&n)
	fmt.Scan(&q)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	var que processQueue = make([]*process, 0)
	for i := 0; i < n; i++ {
		s.Scan()
		name := s.Text()
		s.Scan()
		time, _ := strconv.Atoi(s.Text())
		que.push(&process{name: name, time: time})
	}

	t := 0
	for len(que) > 0 {
		p := que.pop()
		if p.time <= q {
			t += p.time
			fmt.Printf("%s %d\n", p.name, t)
		} else {
			t += q
			p.time -= q
			que.push(p)
		}
	}
}
