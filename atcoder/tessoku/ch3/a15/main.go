package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := load(s)
	a := make([]int, n)
	set := make(map[int]struct{})
	size := n
	for i := 0; i < n; i++ {
		a[i] = load(s)
		if _, ok := set[a[i]]; ok {
			size--
		} else {
			set[a[i]] = struct{}{}
		}
	}
	keys := make([]int, size)
	i := 0
	for k := range set {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	rank := make(map[int]int)
	for i := 0; i < size; i++ {
		rank[keys[i]] = i + 1
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprint(w, rank[a[0]])
	for i := 1; i < n; i++ {
		fmt.Fprint(w, " ")
		fmt.Fprint(w, rank[a[i]])
	}
	fmt.Fprintln(w)
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
