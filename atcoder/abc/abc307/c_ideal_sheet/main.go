package main

import (
	"fmt"
	"strings"
)

func main() {
	ha, wa, a := scan()
	hb, wb, b := scan()
	hx, wx, x := scan()

	for ia := -ha+1; ia < hx; ia++ {
		for ja := -wa+1; ja < wx; ja++ {
			for ib := -hb+1; ib < hx; ib++ {
				for jb := -wb+1; jb < wx; jb++ {
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
