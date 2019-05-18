package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var h int
	fmt.Scan(&h)

	heap := make([]int, h+1)
	
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 1; i <= h; i++ {
		sc.Scan()
		heap[i], _ = strconv.Atoi(sc.Text())
	}
	for i := h/2; i >= 1; i-- {
		maxHeapify(&heap, i)
	}

	wr := bufio.NewWriter(os.Stdout)
	for i := 1; i <= h; i++ {
		wr.WriteString(" ")
		wr.WriteString(strconv.Itoa(heap[i]))
	}
	wr.WriteString("\n")
	wr.Flush()
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
