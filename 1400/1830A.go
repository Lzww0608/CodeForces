package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	type pair struct {
		x, i int
	}
	var n int
	fmt.Fscan(in, &n)
	g := make([][]pair, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], pair{v, i})
		g[v] = append(g[v], pair{u, i})
	}

	a := make([]int, n)
	f := make([]int, n)
	a[0] = n
	var dfs func(int, int)
	dfs = func(u, fa int) {
		for _, v := range g[u] {
			if v.x == fa {
				continue
			}
			a[v.x] = v.i
			if v.i < a[u] {
				f[v.x] = f[u] + 1
			} else {
				f[v.x] = f[u]
			}
			dfs(v.x, u)
		}
	}
	dfs(0, -1)

	fmt.Fprintln(out, slices.Max(f))
}
