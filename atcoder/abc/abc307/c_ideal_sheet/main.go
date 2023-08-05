package main

import (
	"fmt"
	"strings"
)

func main() {
	ha, wa, a := scan()
	hb, wb, b := scan()
	hx, wx, x := scan()

	for ia := 0; ia+ha <= hx; ia++ {
		for ja := 0; ja+wa <= wx; ja++ {
			for ib := 0; ib+hb <= hx; ib++ {
				for jb := 0; jb+wb <= wx; jb++ {
					ab := make([]string, hx)
					for i := 0; i < hx; i++ {
						ab[i] = strings.Repeat(".", wx)
						for j := 0; j < wx; j++ {

						}
					}
				}
			}
		}
	}
}

func scan() (h, w int, s []string) {
	fmt.Scan(&h)
	fmt.Scan(&w)
	s = make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}
	return
}
