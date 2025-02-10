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

	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]int, n)
	rg := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		rg[v] = append(rg[v], u)
	}

	dis := make([][]int, n)
	rdis := make([][]int, n)

	bfs := func(start int) {
		q := []int{start}
		vis := make([]bool, n)
		vis[start] = true
		d := 0
		for len(q) > 0 {
			for sz := len(q); sz > 0; sz-- {
				u := q[0]
				q = q[1:]
				dis[start][u] = d
				for _, v := range g[u] {
					if !vis[v] {
						vis[v] = true
						q = append(q, v)
					}
				}
			}
			d++
		}
	}
	rbfs := func(start int) {
		q := []int{start}
		vis := make([]bool, n)
		vis[start] = true
		d := 0
		for len(q) > 0 {
			for sz := len(q); sz > 0; sz-- {
				u := q[0]
				q = q[1:]
				rdis[start][u] = d
				for _, v := range rg[u] {
					if !vis[v] {
						vis[v] = true
						q = append(q, v)
					}
				}
			}
			d++
		}
	}

	d := make([][][2]int, n)
	rd := make([][][2]int, n)
	id := make([]int, n)
	for j := 0; j < n; j++ {
		id[j] = j
	}
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		rdis[i] = make([]int, n)
		bfs(i)
		rbfs(i)
		sort.Slice(id, func(p, q int) bool {
			return dis[i][id[p]] > dis[i][id[q]]
		})
		for j := 0; j < 3 && j < len(dis[i]); j++ {
			d[i] = append(d[i], [2]int{id[j], dis[i][id[j]]})
		}
		sort.Slice(id, func(p, q int) bool {
			return rdis[i][id[p]] > rdis[i][id[q]]
		})
		for j := 0; j < 3 && j < len(dis[i]); j++ {
			rd[i] = append(rd[i], [2]int{id[j], rdis[i][id[j]]})
		}
	}

	ans := [4]int{}
	mx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j || dis[i][j] == 0 {
				continue
			}
			sum := dis[i][j]
			for _, u := range d[j] {
				if u[0] == i || u[0] == j || u[1] == 0 {
					continue
				}
				for _, v := range rd[i] {
					if v[0] == i || v[0] == j || v[0] == u[0] || v[1] == 0 {
						continue
					}
					if sum+v[1]+u[1] > mx {
						mx = sum + v[1] + u[1]
						ans = [4]int{v[0], i, j, u[0]}
					}
				}
			}
		}
	}
	for _, v := range ans {
		fmt.Fprint(out, v+1, " ")
	}
}
