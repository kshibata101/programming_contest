package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type card struct {
	s string
	v int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	cards := make([]card, n)
	for i := 0; i < n; i++ {
		s.Scan()
		cards[i] = card{s: s.Text()}
		s.Scan()
		cards[i].v, _ = strconv.Atoi(s.Text())
	}
	cards2 := make([]card, n)
	copy(cards2, cards)

	quickSort(cards2, 0, n-1)

	// check stable or not
	chars := make(map[int][]string)
	indexs := make(map[int]int)
	for i := 0; i < n; i++ {
		card := cards2[i]
		chars[card.v] = append(chars[card.v], card.s)
		indexs[card.v] = 0
	}
	stable := true
	for i := 0; i < n; i++ {
		card := cards[i]
		index, _ := indexs[card.v]
		strs, _ := chars[card.v]
		if card.s != strs[index] {
			stable = false
			break
		}
		indexs[card.v]++
	}

	// print
	if stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%s %v\n", cards2[i].s, cards2[i].v)
	}
}

func quickSort(a []card, p, r int) {
	if p >= r {
		return
	}
	q := partition(a, p, r)
	quickSort(a, p, q-1)
	quickSort(a, q+1, r)
}

func partition(a []card, p, r int) int {
	i := p
	for j := p; j < r; j++ {
		if a[j].v <= a[r].v {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}
