package main

import (
	"fmt"
	"strings"
)

func main() {
	ha, wa, a := scan()
	hb, wb, b := scan()
	hx, wx, x := scan()

	sharp := []rune("#")[0]
	for ia := -ha + 1; ia < ha+hx-1; ia++ {
		for ja := -wa + 1; ja < wa+wx-1; ja++ {
			for ib := -hb + 1; ib < hb+hx-1; ib++ {
				for jb := -wb + 1; jb < wb+wx-1; jb++ {
					success := true
					xx := make([][]rune, hx)
					for ix := 0; ix < hx; ix++ {
						xx[ix] = []rune(strings.Repeat(".", wx))
					}
					// copy a to xx
					for ai := 0; ai < ha; ai++ {
						for aj := 0; aj < wa; aj++ {
							if string(a[ai][aj]) == "#" {
								if 0 <= ai+ia && ai+ia < hx && 0 <= aj+ja && aj+ja < wx {
									xx[ai+ia][aj+ja] = sharp
								} else {
									success = false
									break
								}
							}
						}
						if !success {
							break
						}
					}
					if !success {
						break
					}
					for bi := 0; bi < hb; bi++ {
						for bj := 0; bj < wb; bj++ {
							if string(b[bi][bj]) == "#" {
								if 0 <= bi+ib && bi+ib < hx && 0 <= bj+jb && bj+jb < wx {
									xx[bi+ib][bj+jb] = sharp
								} else {
									success = false
									break
								}
							}
						}
						if !success {
							break
						}
					}
					for ix := 0; ix < hx; ix++ {
						if string(xx[ix]) != x[ix] {
							success = false
							break
						}
					}
					if success {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}
	fmt.Println("No")
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
