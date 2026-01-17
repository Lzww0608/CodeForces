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

	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]int, n)
	for range m {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	color := make([]int, n)
	a := []int{}
	b := []int{}
	var dfs func(int, int) bool
	dfs = func(u, c int) bool {
		if color[u] != 0 {
			if color[u] == c {
				return true
			}
			return false
		}

		color[u] = c
		if c == 1 {
			a = append(a, u+1)
		} else {
			b = append(b, u+1)
		}
		for _, v := range g[u] {
			if !dfs(v, 3-c) {
				return false
			}
		}

		return true
	}

	for i := range n {
		if color[i] == 0 {
			if !dfs(i, 1) {
				fmt.Fprintln(out, -1)
				return
			}
		}
	}

	fmt.Fprintln(out, len(a))
	for _, x := range a {
		fmt.Fprintf(out, "%d ", x)
	}
	fmt.Fprintln(out)
	fmt.Fprintln(out, len(b))
	for _, x := range b {
		fmt.Fprintf(out, "%d ", x)
	}
}
