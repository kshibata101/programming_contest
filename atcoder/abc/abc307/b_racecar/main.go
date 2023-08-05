package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	t := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			str := t[i] + t[j]
			l := len(str)
			kaibun := true
			for k := 0; k < (l+1)/2; k++ {
				if str[k] != str[l-k-1] {
					kaibun = false
					break
				}
			}
			if kaibun {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
