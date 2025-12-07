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
	type pair struct {
		l, r int
	}
	a := make([]pair, n)
	for i := range a {
		fmt.Fscan(in, &a[i].l, &a[i].r)
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

	var dfs func(int, int) (int, int)
	dfs = func(u, fa int) (int, int) {
		x, y := 0, 0
		for _, v := range g[u] {
			if v == fa {
				continue
			}

			t1, t2 := dfs(v, u)
			// l
			x += max(t1+abs(a[u].l-a[v].l), t2+abs(a[u].l-a[v].r))
			y += max(t1+abs(a[u].r-a[v].l), t2+abs(a[u].r-a[v].r))
		}

		return x, y
	}

	fmt.Fprintln(out, max(dfs(0, -1)))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
