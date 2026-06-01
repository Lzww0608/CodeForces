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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
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

	ans := make([]int, n)
	var dfs func(int, int) int
	dfs = func(u, fa int) int {
		cnt := 1
		if a[u] == 0 {
			cnt = -1
		}
		for _, v := range g[u] {
			if v != fa {
				cnt += max(0, dfs(v, u))
			}
		}
		ans[u] = cnt
		return cnt
	}
	dfs(0, -1)

	var dfs2 func(int, int)
	dfs2 = func(u, fa int) {
		if fa != -1 {
			if ans[u] > 0 {
				ans[u] += max(0, ans[fa]-ans[u])
			} else {
				ans[u] += max(0, ans[fa])
			}
		}

		for _, v := range g[u] {
			if v != fa {
				dfs2(v, u)
			}

		}
	}
	dfs2(0, -1)

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}
