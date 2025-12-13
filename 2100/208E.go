package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 17

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, x int
	fmt.Fscan(in, &n)
	g := make([][]int, n+1)
	var roots []int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x)
		if x == 0 {
			roots = append(roots, i)
		} else {
			g[x] = append(g[x], i)
			g[i] = append(g[i], x)
		}
	}

	t := 0
	type info struct {
		tin, tout, d int
	}
	f := make([][N]int, n+1)
	a := make([]info, n+1)
	dep := make([][]int, n+1)
	var dfs func(u, fa, d int)
	dfs = func(u, fa, d int) {
		t++
		f[u][0] = fa
		a[u] = info{
			tin: t,
			d:   d,
		}
		dep[d] = append(dep[d], t)
		for _, v := range g[u] {
			if v != fa {
				dfs(v, u, d+1)
			}
		}

		a[u].tout = t
	}
	for _, root := range roots {
		dfs(root, -1, 0)
	}

	// binary lifting
	for i := range N - 1 {
		for j := 1; j <= n; j++ {
			if p := f[j][i]; p != -1 {
				f[j][i+1] = f[p][i]
			} else {
				f[j][i+1] = -1
			}
		}
	}

	findKthAncestor := func(v, k int) int {
		for i := range N {
			if (k>>i)&1 == 1 {
				v = f[v][i]
			}
			if v == -1 {
				break
			}
		}

		return v
	}

	query := func(v, k int) int {
		k += a[v].d

		l := sort.SearchInts(dep[k], a[v].tin)
		r := sort.SearchInts(dep[k], a[v].tout+1)
		return r - l - 1
	}

	fmt.Fscan(in, &m)
	ans := make([]int, m)
	var v, p int
	for i := range m {
		fmt.Fscan(in, &v, &p)
		v = findKthAncestor(v, p)
		if v != -1 {
			ans[i] = query(v, p)
		}
	}

	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
}
