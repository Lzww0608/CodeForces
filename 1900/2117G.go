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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]pair, n)
	f := make([][3]int, 0, m)
	var u, v, w int
	for range m {
		fmt.Fscan(in, &u, &v, &w)
		u--
		v--
		f = append(f, [3]int{u, v, w})
		g[u] = append(g[u], pair{v, w})
		g[v] = append(g[v], pair{u, w})
	}

	dij := func(start int) []int {
		dis := make([]int, n)
		vis := make([]bool, n)
		for i := range dis {
			dis[i] = INF
		}
		dis[start] = 0
		h := &hp{}
		heap.Init(h)
		heap.Push(h, pair{start, 0})
		for h.Len() > 0 {
			p := heap.Pop(h).(pair)
			if vis[p.u] {
				continue
			}
			vis[p.u] = true
			for _, t := range g[p.u] {
				if vis[t.u] {
					continue
				}
				d := max(t.w, p.w)
				if dis[t.u] > d {
					dis[t.u] = d
					heap.Push(h, pair{t.u, d})
				}
			}
		}

		return dis
	}
	dis1 := dij(0)
	disn := dij(n - 1)

	ans := INF
	for _, t := range f {
		ans = min(ans, t[2]+max(dis1[t[0]], disn[t[1]], t[2]))
		ans = min(ans, t[2]+max(dis1[t[1]], disn[t[0]], t[2]))
	}

	fmt.Fprintln(out, ans)
}

const INF = math.MaxInt64 / 2

type pair struct {
	u, w int
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].w < h[j].w }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() (v interface{}) {
	n := len(*h)
	v = (*h)[n-1]
	*h = (*h)[:n-1]
	return v
}
