package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	ans := 0
	for i := 1; i <= k/3 && i <= n; i++ {
		for j := i; j <= k/2 && j <= n; j++ {
			l := k - i - j
			if l < j {
				break
			}
			if l > n {
				continue
			}
			if i == j {
				if j == l {
					ans += 1
				} else {
					ans += 3
				}
			} else {
				if j == l {
					ans += 3
				} else {
					ans += 6
				}
			}
		}
	}
	fmt.Println(ans)
}