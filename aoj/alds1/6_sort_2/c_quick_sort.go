package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Mark string
	Num int
}

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	cards := make([]Card, n)
	cards2 := make([]Card, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		arr := strings.Split(sc.Text(), " ")
		num, _ := strconv.Atoi(arr[1])
		cards[i] = Card{arr[0], num}
		cards2[i] = Card{arr[0], num}
	}

	quickSort(&cards, 0, len(cards)-1)
	mergeSort(&cards2, 0, len(cards))

	stable := true
	for i := 0; i < n; i++ {
		if cards[i].Mark != cards2[i].Mark {
			stable = false
			break
		}
	}

	if stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}

	for i := 0; i < n; i++ {
		fmt.Println(cards[i].Mark, cards[i].Num)
	}
}

func quickSort(a *[]Card, p, r int) {
	if p < r {
		q := partition(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q+1, r)
	}
}

func partition(a *[]Card, p, r int) int {
	x := (*a)[r].Num
	i := p
	for j := p; j < r; j++ {
		if (*a)[j].Num <= x {
			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
			i++
		}
	}
	(*a)[i], (*a)[r] = (*a)[r], (*a)[i]
	return i
}

func mergeSort(a *[]Card, l, r int) {
	if l + 1 < r {
		m := (l + r) / 2
		mergeSort(a, l, m)
		mergeSort(a, m, r)
		merge(a, l, m, r)
	}
}

func merge(a *[]Card, l, m, r int) {
	ll := append([]Card{}, (*a)[l:m]...)
	rr := append([]Card{}, (*a)[m:r]...)

	// 番兵
	ll = append(ll, Card{"-", 1 << 30})
	rr = append(rr, Card{"-", 1 << 30})

	i, j := 0, 0
	for k := l; k < r; k++ {
		if ll[i].Num <= rr[j].Num {
			(*a)[k] = ll[i]
			i++
		} else {
			(*a)[k] = rr[j]
			j++
		}
	}
}
