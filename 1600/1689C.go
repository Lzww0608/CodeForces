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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	g := make([][]int, n)

	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	ans := 0
	dep := make([]int, n)
	var dfs func(int, int)
	dfs = func(u, fa int) {
		deg := 0
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			dep[v] = dep[u] + 1
			deg++
			dfs(v, u)
		}

		if deg < 2 {
			ans = max(ans, n-2*dep[u]-1-deg)
		}
	}

	dfs(0, -1)
	fmt.Fprintln(out, ans)
}
