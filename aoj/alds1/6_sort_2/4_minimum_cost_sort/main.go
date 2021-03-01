package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	w := make([]int, n)
	done := make([]bool, 10001)

	min := 10000
	for i := 0; i < n; i++ {
		s.Scan()
		w[i], _ = strconv.Atoi(s.Text())
		if w[i] < min {
			min = w[i]
		}
	}
	loop := buildLoop(w, n)

	ans := 0
	for i := 0; i < n; i++ {
		if done[w[i]] {
			continue
		}
		sum := 0
		cnt := 0
		minInLoop := 10000
		now := w[i]
		for !done[now] {
			done[now] = true
			sum += now
			cnt++
			if now < minInLoop {
				minInLoop = now
			}
			now = loop[now]
		}
		ans1 := sum + minInLoop*(cnt-2)
		ans2 := sum + minInLoop + min*(cnt+1)
		if ans1 < ans2 {
			ans += ans1
		} else {
			ans += ans2
		}
	}
	fmt.Println(ans)
}

func buildLoop(w []int, n int) map[int]int {
	sw := make([]int, n)
	copy(sw, w)
	quickSort(sw, 0, n-1)
	loop := make(map[int]int)
	for i := 0; i < n; i++ {
		loop[w[i]] = sw[i]
	}
	return loop
}

func quickSort(w []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(w, p, r)
	quickSort(w, p, q-1)
	quickSort(w, q+1, r)
}

func partition(w []int, p, r int) int {
	i := p
	for j := p; j < r; j++ {
		if w[j] <= w[r] {
			w[i], w[j] = w[j], w[i]
			i++
		}
	}
	w[i], w[r] = w[r], w[i]
	return i
}
