package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

func ScanS() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	h := Scan()
	w := Scan()
	k := Scan()
	c := make([][]bool, h)
	for i := 0; i < h; i++ {
		c[i] = make([]bool, w)
		s := ScanS()
		for j := 0; j < len(s); j++ {
			if s[j] == '#' {
				c[i][j] = true
			}
		}
	}

	ans := 0
	for i := 0; i < (1 << h); i++ {
		remains := k
		cc := make([][]bool, h)
		for j := 0; j < h; j++ {
			cc[j] = make([]bool, w)
			if (i>>j)&1 > 0 {
				// 行が選択されたので全部塗りつぶす
				for l := 0; l < w; l++ {
					cc[j][l] = true
				}
				remains--
				if remains < 0 {
					break
				}
			} else {
				// 行が選択されなかったので元のマスをコピー
				copy(cc[j], c[j])
			}
		}
		if remains < 0 {
			continue
		}
		// 列ごとに黒の数を数える
		black := make([]int, w)
		num := 0
		for j := 0; j < h; j++ {
			for l := 0; l < w; l++ {
				if cc[j][l] {
					black[l]++
					num++
				}
			}
		}
		// 黒の少ないところ(=白の多いところ)順に黒に変えていく
		sort.Ints(black)
		for j := 0; j < remains; j++ {
			num += h - black[j]
		}
		if num > ans {
			ans = num
		}
	}
	fmt.Println(ans)
}
