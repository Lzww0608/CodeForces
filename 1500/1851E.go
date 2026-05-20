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
	c := make([]int, n)
	for i := range c {
		fmt.Fscan(in, &c[i])
	}
	p := make([]int, k)
	for i := range p {
		fmt.Fscan(in, &p[i])
	}
	e := make([][]int, n)
	for i := range e {
		var m int
		fmt.Fscan(in, &m)
		e[i] = make([]int, m)
		for j := range e[i] {
			fmt.Fscan(in, &e[i][j])
		}
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	for _, x := range p {
		ans[x-1] = 0
	}

	var dfs func(int) int
	dfs = func(u int) int {
		if ans[u] != -1 {
			return ans[u]
		}

		ans[u] = c[u]
		cur := 0

		if len(e[u]) > 0 {
			for _, v := range e[u] {
				cur += dfs(v - 1)
			}
			ans[u] = min(cur, c[u])
		}

		return ans[u]
	}

	for i := range ans {
		ans[i] = dfs(i)
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
