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
	g := make([][][2]int, n)
	for i := 0; i < n-1; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		u--
		v--
		g[u] = append(g[u], [2]int{v, w})
		g[v] = append(g[v], [2]int{u, w})
	}

	check := func(x int) bool {
		for x > 0 {
			t := x % 10
			if t != 4 && t != 7 {
				return false
			}
			x /= 10
		}
		return true
	}

	vis := make([]bool, n)
	var dfs func(int, int) int
	dfs = func(u, fa int) (res int) {
		if vis[u] {
			return 0
		}
		vis[u] = true
		res = 1
		for _, v := range g[u] {
			if v[0] == fa {
				continue
			}
			if !check(v[1]) {
				res += dfs(v[0], u)
			}
		}

		return res
	}

	ans := 0
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		x := dfs(i, -1)
		ans += x * (n - x) * (n - x - 1)
	}
	fmt.Fprintln(out, ans)
}
