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
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
	}

	for range m {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u][v] = true
		g[v][u] = true
	}

	route := !g[0][n-1]
	q := []int{0}
	used := make([]bool, n)
	used[0] = true
	for ans := 1; len(q) > 0; ans++ {
		for t := len(q); t > 0; t-- {
			u := q[0]
			q = q[1:]
			for v := 0; v < n; v++ {
				if g[u][v] == route && !used[v] {
					if v == n-1 {
						fmt.Fprintln(out, ans)
						return
					}
					used[v] = true
					q = append(q, v)
				}
			}
		}
	}

	fmt.Fprintln(out, -1)
}
