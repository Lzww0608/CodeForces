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
	a := make([]int, n)
	sum := 0
	for i := range a {
		fmt.Fscan(in, &a[i])
		sum ^= a[i]
	}
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	if sum == 0 {
		fmt.Fprintln(out, "YES")
		return
	}

	var dfs func(int, int) int
	dfs = func(u, fa int) (ans int) {
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			ans += dfs(v, u)
			a[u] ^= a[v]
		}
		if ans > 0 {
			if a[u] == 0 {
				ans++
			}
		} else if a[u] == sum {
			ans++
		}
		return
	}

	if k > 2 && dfs(0, -1) > 1 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
