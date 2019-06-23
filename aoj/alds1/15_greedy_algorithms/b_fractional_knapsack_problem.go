package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

type Item struct {
	W float64
	V float64
}

type Items []Item

func (is Items) Len() int { return len(is) }

func (is Items) Swap(i,j int) { is[i], is[j] = is[j], is[i]}

func (is Items) Less (i,j int) bool {
	return is[i].V / is[i].W > is[j].V / is[j].W
}

func main() {
	var n int
	fmt.Scan(&n)
	var w float64
	fmt.Scan(&w)
	var items Items = make([]Item, n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		items[i].V, _ = strconv.ParseFloat(sc.Text(), 64)
		sc.Scan()
		items[i].W, _ = strconv.ParseFloat(sc.Text(), 64)
	}
	sort.Sort(items)

	val := 0.0
	for i := 0; i < n; i++ {
		if items[i].W >= w {
			val += items[i].V * w / items[i].W
			break
		}
		val += items[i].V
		w -= items[i].W
	}
	fmt.Printf("%f\n", val)
}