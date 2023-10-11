package main

import "fmt"

func Permute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(nums, &ret)
	return ret
}

func permute(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}

func main() {
	var c []int = make([]int, 9)
	for i := 0; i < len(c); i++ {
		fmt.Scan(&c[i])
	}
	var a []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	perm := Permute(a)

	checkMap := [][]int{
		[]int{0, 3, 6},
		[]int{0, 4},
		[]int{0, 5, 7},
		[]int{1, 3},
		[]int{1, 4, 6, 7},
		[]int{1, 5},
		[]int{2, 3, 7},
		[]int{2, 4},
		[]int{2, 5, 6},
	}
	g := 0
	for i := 0; i < len(perm); i++ {
		check := make([][]int, 8)
		for j := 0; j < 8; j++ {
			check[j] = []int{}
		}

		for j := 0; j < 9; j++ {
			num := perm[i][j]
			for k := 0; k < len(checkMap[num]); k++ {
				index := checkMap[num][k]
				check[index] = append(check[index], c[num])
			}
		}
		gakkari := false
		for j := 0; j < 8; j++ {
			if check[j][0] == check[j][1] && check[j][1] != check[j][2] {
				gakkari = true
				break
			}
		}
		if gakkari {
			g++
		}
	}
	fmt.Println(1 - float64(g)/float64(len(perm)))
}
