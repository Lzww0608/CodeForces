package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)
	fa := make([]int, n+1)
	g := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &fa[i])
		g[fa[i]] = append(g[fa[i]], i)
	}

	cnt := make([]int, n+1)
	sz := make([]int, n+1)

	var f func(int)
	f = func(u int) {
		cnt[u] = u
		sz[u] = 1
		mx := 0
		for _, v := range g[u] {
			f(v)
			sz[u] += sz[v]
			if sz[v] > sz[mx] {
				mx = v
			}
		}

		if sz[mx]*2 > sz[u] {
			x := cnt[mx]
			for sz[x]*2 < sz[u] {
				x = fa[x]
			}
			cnt[u] = x
		}
	}
	f(1)

	var x int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &x)
		fmt.Fprintln(out, cnt[x])
	}
}
