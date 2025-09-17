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
	var n, k int
	fmt.Fscan(in, &n, &k)
	k--
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	var dfs func(int, int) int
	dfs = func(u, fa int) (res int) {
		res = 1
		for _, v := range g[u] {
			if v != fa {
				res += dfs(v, u)
			}
		}

		return
	}

	if len(g[k]) <= 1 {
		fmt.Fprintln(out, "Ayush")
		return
	}

	cnt := dfs(k, -1) - 1
	if cnt&1 == 1 {
		fmt.Fprintln(out, "Ayush")
	} else {
		fmt.Fprintln(out, "Ashish")
	}
}
