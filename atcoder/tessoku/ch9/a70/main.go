package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func Arr2Bin(a []int) int {
	ans := 0
	for i := 0; i < len(a); i++ {
		ans += a[i] * (1 << i)
	}
	return ans
}

func Bin2Arr(a, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = a % 2
		a /= 2
	}
	return arr
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	m := Scan()
	a := make([]int, n)
	for j := 0; j < n; j++ {
		a[j] = Scan()
	}
	s := Arr2Bin(a)

	x := make([]int, m)
	y := make([]int, m)
	z := make([]int, m)
	for j := 0; j < m; j++ {
		x[j] = Scan() - 1
		y[j] = Scan() - 1
		z[j] = Scan() - 1
	}

	G := make([][]int, 1<<n)
	for i := 0; i < (1 << n); i++ {
		G[i] = make([]int, 0, m)
	}
	for i := 0; i < (1 << n); i++ {
		for j := 0; j < m; j++ {
			arr := Bin2Arr(i, n)
			arr[x[j]] = arr[x[j]] ^ 1
			arr[y[j]] = arr[y[j]] ^ 1
			arr[z[j]] = arr[z[j]] ^ 1
			G[i] = append(G[i], Arr2Bin(arr))
		}
	}

	q := []int{s}
	dist := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dist[i] = -1
	}
	dist[s] = 0

	for len(q) > 0 {
		// pop
		v := q[0]
		if v == (1<<n)-1 {
			fmt.Println(dist[v])
			return
		}
		if len(q) == 1 {
			q = []int{}
		} else {
			q = q[1:]
		}

		for i := 0; i < len(G[v]); i++ {
			next := G[v][i]
			if dist[next] == -1 {
				q = append(q, next)
				dist[next] = dist[v] + 1
			}
		}

	}
	fmt.Println(-1)
}
