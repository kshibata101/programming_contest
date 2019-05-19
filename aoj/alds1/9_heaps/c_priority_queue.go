package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	pq := []int{-1}
	
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr := bufio.NewWriter(os.Stdout)
	for {
		sc.Scan()
		c := sc.Text()
		if c == "end" {
			break
		} else if c ==  "extract" {
			m := extractMax(&pq)
			wr.WriteString(strconv.Itoa(m))
			wr.WriteString("\n")
		} else if c == "insert" {
			sc.Scan()
			k, _ := strconv.Atoi(sc.Text())
			insert(&pq, k)
		}
	}
	wr.Flush()
}

func insert(pq *[]int, k int) {
	*pq = append(*pq, k)
	for n := len(*pq) - 1; n > 1; n /= 2 {
		if (*pq)[n] <= (*pq)[n/2] {
			break
		}
		(*pq)[n], (*pq)[n/2] = (*pq)[n/2], (*pq)[n]
	}
}

func extractMax(pq *[]int) int {
	max := (*pq)[1]
	if len(*pq) > 2 {
		(*pq)[1] = (*pq)[len(*pq)-1]
		*pq = (*pq)[0:len(*pq)-1]
		maxHeapify(pq, 1)
	} else {
		*pq = []int{-1}
	}
	return max
}

func maxHeapify(heap *[]int, i int) {
	l := 2*i
	r := 2*i+1

	var large int
	if l < len(*heap) && (*heap)[l] > (*heap)[i] {
		large = l
	} else {
		large = i
	}
	if r < len(*heap) && (*heap)[r] > (*heap)[large] {
		large = r
	}

	if large != i {
		(*heap)[large], (*heap)[i] = (*heap)[i], (*heap)[large]
		maxHeapify(heap, large)
	}
}
