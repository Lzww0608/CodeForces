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
	deg := make([]int, n)
	q := make([]int, 0, n)
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
		deg[u]++
		deg[v]++
	}

	time := make([]int, n)
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
			time[i] = 1
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			deg[v]--
			if deg[v] == 1 {
				time[v] = time[u] + 1
				q = append(q, v)
			}
		}
	}

	ans := 0
	for _, t := range time {
		if t > k || t == 0 {
			ans++
		}
	}

	fmt.Fprintln(out, ans)
}
