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
	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)
	a--
	b--
	g := make([][]int, n)
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		deg[x]++
		deg[y]++
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	if a == b {
		fmt.Fprintln(out, "NO")
		return
	}

	q := []int{}
	s := make(map[int]bool)
	for i, x := range deg {
		if x == 1 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		s[u] = true
		for _, v := range g[u] {
			if deg[v]--; deg[v] == 1 {
				q = append(q, v)
			}
		}
	}

	bfs := func(x int) []int {
		dis := make([]int, n)
		vis := make([]bool, n)
		q := []int{x}
		vis[x] = true
		d := 0
		for len(q) > 0 {
			for sz := len(q); sz > 0; sz-- {
				u := q[0]
				q = q[1:]
				dis[u] = d
				for _, v := range g[u] {
					if !vis[v] {
						q = append(q, v)
						vis[v] = true
					}
				}
			}
			d++
		}
		return dis
	}

	disA := bfs(a)
	disB := bfs(b)

	findEntrance := func(x int) int {
		vis := make([]bool, n)
		q := []int{x}
		vis[x] = true
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			if !s[u] {
				return u
			}
			for _, v := range g[u] {
				if !vis[v] {
					q = append(q, v)
					vis[v] = true
				}
			}
		}
		return -1
	}

	entrance := findEntrance(b)
	if !s[b] || disA[entrance] > disB[entrance] {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
