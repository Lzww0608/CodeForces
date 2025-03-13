package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, l, s, t int
	fmt.Fscan(in, &n, &m, &l, &s, &t)
	edges := make([][]int, m)
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		edges[i] = []int{u, v, w}
		g[u][v] = i
		g[v][u] = i
	}

	dijsktra := func(op, source int, f []int) (ans []int) {
		dist := make([]int, n)
		vis := make([]bool, n)
		for i := range dist {
			dist[i] = math.MaxInt32
		}
		dist[source] = 0
		for i := 0; i < n-1; i++ {
			u := -1
			for j := 0; j < n; j++ {
				if !vis[j] && (u == -1 || dist[j] < dist[u]) {
					u = j
				}
			}
			vis[u] = true
			for v := 0; v < n; v++ {
				if !vis[v] && g[u][v] != -1 {
					k := g[u][v]
					if edges[k][2] != 0 {
						dist[v] = min(dist[v], dist[u]+edges[k][2])
					} else if op == 0 {
						dist[v] = min(dist[v], dist[u]+1)
					} else {
						modify := l - dist[u] - f[v]
						if modify > 0 {
							dist[v] = min(dist[v], dist[u]+modify)
							edges[k][2] = modify
						} else {
							edges[k][2] = l
						}
					}
				}
			}
		}

		return dist
	}

	f := dijsktra(0, t, nil)
	if f[s] > l {
		fmt.Fprintln(out, "NO")
		return
	}
	f = dijsktra(1, s, f)
	if f[t] != l {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for _, v := range edges {
		fmt.Fprintln(out, v[0], v[1], v[2])
	}
}
