package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var s *bufio.Scanner

func Scan() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

type Movie struct {
	L int
	R int
}

func search(movies []Movie, t int) int {
	l := 0
	r := len(movies)
	for l < r {
		m := (l + r) / 2
		if t < movies[m].R {
			r = m
		} else if t > movies[m].R {
			l = m + 1
		} else {
			return m
		}
	}
	return l
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := Scan()
	movies := make([]Movie, n)
	for i := 0; i < n; i++ {
		l := Scan()
		r := Scan()
		movies[i] = Movie{L: l, R: r}
	}
	sort.Slice(movies, func(i, j int) bool { return movies[i].R < movies[j].R })
	ans := 0
	t := 0
	for t < 86400 {
		i := search(movies, t)
		for ; i < n; i++ {
			if movies[i].L >= t {
				break
			}
		}
		if i >= n {
			break
		}
		ans++
		t = movies[i].R
	}
	fmt.Println(ans)
}
