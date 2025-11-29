package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]int, n+1)
	for range m {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	ans := make([]int, 0, n)
	ans = append(ans, 1)
	vis := make([]bool, n+1)
	vis[1] = true

	h := &hp{}
	for _, v := range g[1] {
		heap.Push(h, v)
	}

	for h.Len() > 0 {
		u := heap.Pop(h).(int)
		if vis[u] {
			continue
		}
		vis[u] = true
		ans = append(ans, u)
		for _, v := range g[u] {
			if !vis[v] {
				heap.Push(h, v)
			}
		}
	}

	for _, v := range ans {
		fmt.Fprintf(out, "%d ", v)
	}
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
