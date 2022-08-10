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

type queue []*process

func (que *queue) push(p *process) {
	*que = append(*que, p)
}

func (que *queue) pop() (res *process) {
	res = (*que)[0]
	*que = (*que)[1:]
	return
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	q, _ := strconv.Atoi(s.Text())

	var que queue = make([]*process, n)
	for i := 0; i < n; i++ {
		s.Scan()
		n := s.Text()

		s.Scan()
		t, _ := strconv.Atoi(s.Text())
		que[i] = &process{name: n, time: t}
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
