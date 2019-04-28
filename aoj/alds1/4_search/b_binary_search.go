package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func binarySearch(list []int, target int, l int, u int) int {
	if u - l <= 1 {
		if list[u] == target {
			return u
		} else {
			return -1
		}
	}

	m := (u + l) / 2
	if target <= list[m] {
		return binarySearch(list, target, l, m)
	} else {
		return binarySearch(list, target, m, u)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	s := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i], _ = strconv.Atoi(sc.Text())
	}

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	
	cnt := 0
	for i := 0; i < q; i++ {
		sc.Scan()
		target, _ := strconv.Atoi(sc.Text())
		
		result := binarySearch(s, target, -1, n)
		if result >= 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
