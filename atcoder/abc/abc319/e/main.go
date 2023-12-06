package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, X, Y int
	fmt.Scan(&N)
	fmt.Scan(&X)
	fmt.Scan(&Y)

	P := make([]int, N)
	T := make([]int, N)
	L := 1

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 1; i < N; i++ {
		sc.Scan()
		P[i], _ = strconv.Atoi(sc.Text())
		sc.Scan()
		T[i], _ = strconv.Atoi(sc.Text())

		L = lcm(L, P[i])
	}

	f := make([]int, L)
	for i := 0; i < L; i++ {
		now := X + i
		for j := 1; j < N; j++ {
			now = (now+P[j]-1)/P[j]*P[j] + T[j]
		}
		f[i] = now + Y - i
	}

	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < Q; i++ {
		sc.Scan()
		q, _ := strconv.Atoi(sc.Text())
		fmt.Fprintln(w, q+f[q%L])
	}
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

func lcm(a, b int) int {
	if a >= b {
		return a * b / gcd(a, b)
	} else {
		return a * b / gcd(b, a)
	}
}
