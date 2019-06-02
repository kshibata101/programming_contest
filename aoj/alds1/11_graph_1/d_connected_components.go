package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var n,m int
	fmt.Scan(&n)
	group := make([]int, n)
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = []int{}
		group[i] = -1
	}
	
	fmt.Scan(&m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		sc.Scan()
		s, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())
		
		g[s] = append(g[s], t)
		g[t] = append(g[t], s)
	}
	for i := 0; i < n; i++ {
		if group[i] != -1 {
			continue
		}
		q := []int{i}
		for len(q) > 0 {
			pop := q[0]
			if len(q) > 1 {
				q = q[1:]
			} else {
				q = []int{}
			}
			if group[pop] == -1 {
				q = append(q, g[pop]...)
				group[pop] = i
			}
		}
	}
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		s, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())

		if group[s] == group[t] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func root(i int, group *[]int) int {
	if i == (*group)[i] {
		return i
	}
	(*group)[i] = root((*group)[i], group)
	return (*group)[i]
}
