package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	I int
	J int
	A int
}

func main() {
	var n int
	fmt.Scan(&n)

	g := make([][]int, n)
	edges := make([]Edge, 0, n*(n-1)/2)
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&g[i][j])
			if g[i][j] >= 0 {
				edges = append(edges, Edge{i, j, g[i][j]})
			}
		}
		parents[i] = i
	}
	sort.SliceStable(edges, func(i, j int) bool { return edges[i].A >= edges[j].A })

	sum := 0
	num := 0
	for len(edges) > 0 && num < n-1 {
		edge := edges[len(edges)-1]
		edges = edges[:len(edges)-1]

		if same(edge.I, edge.J, &parents) {
			continue
		}
		union(edge.I, edge.J, &parents)
		sum += edge.A
		num += 1
	}
	fmt.Println(sum)
}



func root(id int, parents *[]int) int {
	if id == (*parents)[id] {
		return id
	}
	return root((*parents)[id], parents)
}

func union(i, j int, parents *[]int) {
	(*parents)[root(j, parents)] = (*parents)[root(i, parents)]
}

func same(i, j int, parents *[]int) bool {
	return root(i, parents) == root(j, parents)
}
