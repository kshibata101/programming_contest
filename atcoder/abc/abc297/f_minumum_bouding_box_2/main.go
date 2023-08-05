package main

import "fmt"

func permutation(n, k int) int {
	v := 1
	if k > 0 && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n, k int) int {
	return permutation(n, k) / factorial(k)
}

func main() {
	var h, w, k int
	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&k)

	var ans int64 = 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			for i := 0; i < (1 << 4); i++ {
				nx, ny := w, h
				mul := 1
				if i&1 == 1 {
					nx = x
					mul *= -1
				}
				if i&2 == 2 {
					if nx != w {
						nx = 0
					} else {
						nx = w - x + 1
					}
					mul *= -1
				}
				if i&4 == 4 {
					ny = y
					mul *= -1
				}
				if i&8 == 8 {
					if ny != h {
						ny = 0
					} else {
						ny = h - y + 1
					}
					mul *= -1
				}
				c := int64(combination(nx*ny, k) * mul)
				fmt.Printf("c = %d\n", c)
				ans += c
			}
		}
	}

	fmt.Println(float64(ans) / float64(combination(w*h, k)))
}
