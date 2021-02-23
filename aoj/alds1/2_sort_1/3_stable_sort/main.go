package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type card struct {
	suite  rune
	number int
}

func main() {
	var n int
	fmt.Scan(&n)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	cards := []card{}
	for i := 0; i < n; i++ {
		s.Scan()
		runes := []rune(s.Text())
		num, _ := strconv.Atoi(string(runes[1]))
		cards = append(cards, card{suite: runes[0], number: num})
	}
	cards2 := make([]card, n)
	copy(cards2, cards)

	bubbleSort(cards)
	selectionSort(cards2)

	printCards(cards)
	fmt.Println("Stable")

	printCards(cards2)

	stable := true
	for i := 0; i < n; i++ {
		if cards[i].suite != cards2[i].suite ||
			cards[i].number != cards2[i].number {
			stable = false
			break
		}
	}
	if stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}

}

func bubbleSort(cards []card) []card {
	n := len(cards)
	flag := true
	for flag {
		flag = false
		for j := n - 1; j > 0; j-- {
			if cards[j-1].number > cards[j].number {
				cards[j], cards[j-1] = cards[j-1], cards[j]
				flag = true
			}
		}
	}
	return cards
}

func selectionSort(cards []card) []card {
	n := len(cards)
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if cards[j].number < cards[minj].number {
				minj = j
			}
		}
		cards[i], cards[minj] = cards[minj], cards[i]
	}
	return cards
}

func printCards(cards []card) {
	printCard(&cards[0])
	for i := 1; i < len(cards); i++ {
		fmt.Print(" ")
		printCard(&cards[i])
	}
	fmt.Println()
}

func printCard(c *card) {
	fmt.Print(string(c.suite))
	fmt.Print(c.number)
}
