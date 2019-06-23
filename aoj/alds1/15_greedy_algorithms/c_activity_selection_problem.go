package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

type Time struct {
	S int
	T int
}

type Times []Time

func (ts Times) Len() int { return len(ts) }
func (ts Times) Swap(i,j int) { ts[i], ts[j] = ts[j], ts[i] }
func (ts Times) Less(i,j int) bool { return ts[i].T < ts[j].T }

func main() {
	var n int
	fmt.Scan(&n)
	var times Times = make([]Time, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		times[i].S, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		times[i].T, _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(times)

	t, num := 0, 0
	for i := 0; i < n; i++ {
		if t >= times[i].S {
			continue
		}
		t = times[i].T
		num += 1
	}
	fmt.Println(num)
}