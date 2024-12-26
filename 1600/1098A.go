package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	fmt.Fscan(in, &n)
	g := make([][]int, n+1)
	s := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &x)
		g[x] = append(g[x], i)
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &s[i])
	}
	ans := 0

	var dfs func(int, int)
	dfs = func(pre, u int) {
		if s[u] != -1 {
			if s[u] < pre {
				ans = -math.MaxInt64
				return
			}
			ans += s[u] - pre
		} else {
			if len(g[u]) == 0 {
				return
			}
			t := math.MaxInt64
			for _, v := range g[u] {
				t = min(t, s[v]-pre)
			}
			if t < 0 {
				ans = -math.MaxInt64
				return
			}
			ans += t
			pre += t
		}
		for _, v := range g[u] {
			dfs(max(pre, s[u]), v)
		}
	}
	dfs(0, 1)
	if ans < 0 {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, ans)
}
