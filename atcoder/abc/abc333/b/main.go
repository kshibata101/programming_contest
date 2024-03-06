package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	nexts := []string{"AB", "AE", "BA", "BC", "CB", "CD", "DC", "DE", "EA", "ED"}

	ss := 0
	for i := 0; i < len(nexts); i++ {
		if s == nexts[i] { 
			ss = 1
			break
		}
	}
	tt := 0
	for i := 0; i < len(nexts); i++ {
		if t == nexts[i] {
			tt = 1
			break
		}
	}
	if ss ^ tt == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}