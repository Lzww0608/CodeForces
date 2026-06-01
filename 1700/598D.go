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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]string, n)
	g := make([][]int, n)
	res := []int{}
	for i := 0; i < n; i++ {
		g[i] = make([]int, m)
		fmt.Fscan(in, &a[i])
	}

	var dfs func(int, int, int) int
	dfs = func(x, y, id int) int {
		cnt := 0
		q := [][2]int{{x, y}}
		g[x][y] = id
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nx, ny := cur[0]+d[0], cur[1]+d[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= m || g[nx][ny] != 0 {
					continue
				}
				if a[nx][ny] == '*' {
					cnt++
					continue
				}
				g[nx][ny] = id
				q = append(q, [2]int{nx, ny})
			}
		}

		return cnt
	}

	id := 1
	for i := range n {
		for j := range m {
			if a[i][j] == '*' || g[i][j] != 0 {
				continue
			}

			cnt := dfs(i, j, id)
			res = append(res, cnt)
			id++
		}

	}

	for range k {
		var x, y int
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, res[g[x-1][y-1]-1])
	}

}
