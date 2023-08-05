package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	b := 0
	r1, k, r2 := -1, -1, -1
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char == "B" {
			b += i
		}
		if char == "R" {
			if r1 == -1 {
				r1 = i
			} else {
				r2 = i
			}
		}
		if char == "K" {
			k = i
		}
	}
	if b%2 == 1 && r1 < k && k < r2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
