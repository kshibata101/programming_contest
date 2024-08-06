package main

import (
	"fmt"
)

func main() {
	var n int
	var c, s string
	fmt.Scan(&n)
	fmt.Scan(&c)
	fmt.Scan(&s)
	cnt := 0
	for i := 0; i < n; i++ {
		if string(s[i]) == "B" {
			cnt += 1
		} else if string(s[i]) == "R" {
			cnt += 2
		}
	}
	ex := "W"
	if cnt%3 == 1 {
		ex = "B"
	} else if cnt%3 == 2 {
		ex = "R"
	}
	if ex == c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
