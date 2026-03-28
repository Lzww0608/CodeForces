package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	g := make([][]int, n)
	deg := make([]int, n)

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
		deg[u]++
		deg[v]++
	}

	cnt := make([]int, n+1)
	h := &maxHeap{}
	heap.Init(h)

	add := func(x int) {
		if cnt[x] == 0 {
			heap.Push(h, x)
		}
		cnt[x]++
	}

	remove := func(x int) {
		cnt[x]--
	}

	getMaxDeg := func() int {
		for h.Len() > 0 && cnt[(*h)[0]] == 0 {
			heap.Pop(h)
		}
		if h.Len() == 0 {
			return 0
		}
		return (*h)[0]
	}

	for _, x := range deg {
		add(x)
	}

	ans := 0
	for i := 0; i < n; i++ {
		d := deg[i]
		remove(d)

		for _, j := range g[i] {
			remove(deg[j])
			add(deg[j] - 1)
		}

		cur := d + getMaxDeg() - 1
		ans = max(ans, cur)

		for _, j := range g[i] {
			remove(deg[j] - 1)
			add(deg[j])
		}

		add(d)
	}

	fmt.Fprintln(out, ans)
}
