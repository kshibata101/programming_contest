package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	ok := false
	for i := 0; i < n-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			ok = true
			break
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
