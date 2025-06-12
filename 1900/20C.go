package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][][2]int, n+1)
	var u, v, w int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &u, &v, &w)
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, w})
	}

	dis := make([]int, n+1)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[1] = 0
	h := &hp{{1, 0}}

	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)
		if cur[1] > dis[cur[0]] {
			continue
		}

		u, w := cur[0], cur[1]
		for _, v := range g[u] {
			if w+v[1] < dis[v[0]] {
				dis[v[0]] = w + v[1]
				heap.Push(h, [2]int{v[0], dis[v[0]]})
			}
		}
	}

	if dis[n] == math.MaxInt {
		fmt.Fprintln(out, -1)
		return
	}

	ans := make([]int, 0, n)
	ans = append(ans, n)
	cur := n
	for cur != 1 {
		for _, v := range g[cur] {
			if dis[v[0]]+v[1] == dis[cur] {
				cur = v[0]
				ans = append(ans, cur)
			}
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintf(out, "%d ", ans[i])
	}
}

type hp [][2]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.([2]int))
}
func (h *hp) Pop() (v interface{}) {
	n := len(*h)
	v = (*h)[n-1]
	*h = (*h)[:n-1]
	return v
}
