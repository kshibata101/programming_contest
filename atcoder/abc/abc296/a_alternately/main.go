package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	prev := ""
	r := true
	for i := 0; i < n; i++ {
		if prev == string(s[i]) {
			r = false
			break
		}
		prev = string(s[i])
	}
	if r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}