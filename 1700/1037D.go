package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	a := make([]int, n)
	pos := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
		pos[a[i]] = i
	}
	if a[0] != 0 {
		fmt.Fprintln(out, "No")
		return
	}

	for k := range n {
		sort.Slice(g[k], func(i, j int) bool {
			return pos[g[k][i]] < pos[g[k][j]]
		})
	}

	vis := make([]bool, n)
	q := []int{0}
	p := make([]int, 0, n)
	vis[0] = true
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		p = append(p, u)
		for _, v := range g[u] {
			if !vis[v] {
				vis[v] = true
				q = append(q, v)
			}
		}
	}
	//fmt.Fprintln(out, g[1], p)
	for i := range n {
		if p[i] != a[i] {
			fmt.Fprintln(out, "No")
			return
		}
	}

	fmt.Fprintln(out, "Yes")
}
