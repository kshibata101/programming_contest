package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h)
	fmt.Scan(&w)

	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		r := []rune(s)
		prevT := false
		for j := 0; j < w; j++ {
			if r[j] == 'T' {
				if prevT {
					r[j-1] = 'P'
					r[j] = 'C'
					prevT = false
				} else {
					prevT = true
				}
			} else {
				prevT = false
			}
		}
		fmt.Println(string(r))
	}
}
