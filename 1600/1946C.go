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
	g := make([][]int, n)
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	check := func(mid int) bool {
		ans := 0
		var dfs func(int, int) int
		dfs = func(u, p int) int {
			res := 1
			for _, v := range g[u] {
				if v == p {
					continue
				}
				res += dfs(v, u)
			}
			if res >= mid {
				res = 0
				ans++
			}
			return res
		}
		dfs(0, -1)
		return ans > k
	}

	l, r := -1, n
	for l+1 < r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, l)
}
