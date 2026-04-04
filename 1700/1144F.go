package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct {
	u, v int
}

type pair struct {
	to int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	edges := make([]edge, m)
	g := make([][]pair, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		edges[i] = edge{u, v}
		g[u] = append(g[u], pair{v})
		g[v] = append(g[v], pair{u})
	}

	color := make([]int, n)
	for i := range color {
		color[i] = -1
	}

	queue := make([]int, 0, n)
	queue = append(queue, 0)
	color[0] = 0

	for head := 0; head < len(queue); head++ {
		u := queue[head]
		for _, e := range g[u] {
			v := e.to
			if color[v] == -1 {
				color[v] = color[u] ^ 1
				queue = append(queue, v)
				continue
			}
			if color[v] == color[u] {
				fmt.Fprintln(out, "NO")
				return
			}
		}
	}

	ans := make([]byte, m)
	for i, e := range edges {
		if color[e.u] == 0 {
			ans[i] = '0'
		} else {
			ans[i] = '1'
		}
	}

	fmt.Fprintln(out, "YES")
	fmt.Fprintln(out, string(ans))
}
