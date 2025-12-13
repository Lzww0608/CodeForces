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

	var n int
	fmt.Fscan(in, &n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}
	g := make([][]int, n)
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	hson := make([]int, n)
	sz := make([]int, n)
	var dfs func(int, int)
	dfs = func(u, fa int) {
		sz[u] = 1
		mx, x := 0, 0
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			dfs(v, u)
			sz[u] += sz[v]
			if sz[v] > mx {
				mx, x = sz[v], v
			}
		}
		hson[u] = x
	}
	dfs(0, -1)

	ans := make([]any, n)
	var f func(int, int) (map[int]int, map[int]int, int)
	f = func(u, fa int) (map[int]int, map[int]int, int) {
		if hson[u] == 0 {
			ans[u] = c[u]
			return map[int]int{c[u]: 1}, map[int]int{1: c[u]}, 1
		}

		cnt, val, mx := f(hson[u], u)
		merge := func(a, b int) {
			if cnt[a] > 0 {
				val[cnt[a]] -= a
			}

			cnt[a] += b
			val[cnt[a]] += a
			if cnt[a] > mx {
				mx = cnt[a]
			}
		}

		for _, v := range g[u] {
			if v == fa || v == hson[u] {
				continue
			}

			cur, _, _ := f(v, u)
			for a, b := range cur {
				merge(a, b)
			}
		}

		merge(c[u], 1)
		ans[u] = val[mx]
		return cnt, val, mx
	}
	f(0, -1)

	fmt.Fprintln(out, ans...)
}
