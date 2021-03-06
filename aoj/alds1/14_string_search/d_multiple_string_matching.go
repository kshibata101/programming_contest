package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var t string
	fmt.Scan(&t)
	ts := []rune(t)

	// construct suffix array
	ss := makeSuffixArray(t)

	var q int
	fmt.Scan(&q)
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		sc.Scan()
		p := sc.Text()
		ps := []rune(p)

		l := 0
		r := len(ts)
		for r-l > 1 {
			m := (r+l)/2
			idx := int(ss[m+1])
			lst := idx + len(ps)
			if lst >= len(ts) {
				lst = len(ts)
			}
			if p < string(ts[idx:lst]) {
				r = m
			} else {
				l = m
			}
		}
		idx := int(ss[l+1])
		lst := idx + len(ps)
		if lst >= len(ts) {
			lst = len(ts)
		}
		if p == string(ts[idx:lst]) {
			wr.WriteString(strconv.Itoa(1))
		} else {
			wr.WriteString(strconv.Itoa(0))
		}
		wr.WriteString("\n")
	}
	wr.Flush()
}

var n, k int
var ranks []int

type Suffix int

type Suffixes []Suffix

func (ss Suffixes) Len() int { return len(ss)}

func (ss Suffixes) Less(i,j int) bool { return compare(ss[i], ss[j]) }

func (ss Suffixes) Swap(i,j int) { ss[i], ss[j] = ss[j], ss[i] }

func compare(si, sj Suffix) bool {
	oi, oj := int(si), int(sj)
	if ranks[oi] != ranks[oj] {
		return ranks[oi] < ranks[oj]
	}
	ri, rj := -1, -1
	if oi + k <= n {
		ri = ranks[oi+k]
	}
	if oj + k <= n {
		rj = ranks[oj+k]
	}
	return ri < rj
}

func makeSuffixArray(t string) Suffixes {
	rs := []rune(t)
	n = len(rs)
	ranks = make([]int, n+1)
	var ss Suffixes = make([]Suffix, n+1)
	for i := 0; i <= n; i++ {
		ss[i] = Suffix(i)
		if i < n {
			ranks[i] = int(rs[i])
		} else {
			ranks[i] = -1
		}
	}
	k = 1
	for ; k <= n; k*=2 {
		sort.Sort(ss)

		tmp := make([]int, n+1)
		tmp[int(ss[0])] = 0
		for i := 1; i <= n; i++ {
			tmp[int(ss[i])] = tmp[int(ss[i-1])]
			if compare(ss[i-1], ss[i]) {
				tmp[int(ss[i])] += 1
			}
		}
		for i := 0; i <= n; i++ {
			ranks[i] = tmp[i]
		}
	}
	return ss
}