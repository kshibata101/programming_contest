package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type item struct {
	d int // dist
	r int // room
}
type itemHeap []item

func (h itemHeap) Len() int           { return len(h) }
func (h itemHeap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h itemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *itemHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *itemHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	able := make([][]int, n)
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = 1 << 30
		}
	}

	for i := 0; i < m; i++ {
		u := get(sc) - 1
		v := get(sc) - 1
		w := get(sc)
		able[u] = append(able[u], v)
		able[v] = append(able[v], u)
		dis[u][v] = w
		dis[v][u] = w
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	pq := &itemHeap{}
	heap.Init(pq)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = 1 << 30
	}

	// 0日目感染者
	k := get(sc)

	inf := []int{} // 新規感染者リスト
	for i := 0; i < k; i++ {
		r := get(sc) - 1
		inf = append(inf, r)
		ans[r] = 0
	}

	d := get(sc)
	for i := 1; i <= d; i++ {
		x := get(sc)
		for j := 0; j < len(inf); j++ {
			r := inf[j]
			dist[r] = 0
			pq.Push(item{0, r})
		}
		inf = []int{}

		for pq.Len() > 0 {
			it := (*pq)[0]
			if dist[it.r] > x { // 新規感染者から最も近い部屋がその日の感染距離を超えてる場合
				break
			}
			it = pq.Pop().(item)
			if it.d > dist[it.r] {
				continue
			}
			if ans[it.r] == -1 {
				inf = append(inf, it.r)
				ans[it.r] = i
				// ここではdistの更新はしない.してしまうと次の飛び先の計算がit.rまでに来るのに要した距離を無視してしまうため
			}

			from := it.r
			for j := 0; j < len(able[from]); j++ {
				to := able[from][j]
				dj := dis[from][to] + dist[from]
				if dj < dist[to] {
					dist[to] = dj
					pq.Push(item{dj, to})
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}

func get(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}
