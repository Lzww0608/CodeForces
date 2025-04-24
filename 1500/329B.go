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

	var m, n, ans int
	fmt.Fscan(in, &m, &n)
	s := make([]string, m)
	dis := make([][]int, m)

	var sx, sy, ex, ey int
	for i := range s {
		dis[i] = make([]int, n)
		fmt.Fscan(in, &s[i])
		for j := range s[i] {
			if s[i][j] == 'S' {
				sx, sy = i, j
			} else if s[i][j] == 'E' {
				ex, ey = i, j
			}
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cnt := make(map[int]int)
	dis[ex][ey] = 1
	q := [][]int{{ex, ey}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		i, j := cur[0], cur[1]
		if s[i][j] >= '0' && s[i][j] <= '9' {
			cnt[dis[i][j]] += int(s[i][j] - '0')
		}
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || s[x][y] == 'T' || dis[x][y] != 0 {
				continue
			}
			dis[x][y] = dis[i][j] + 1
			q = append(q, []int{x, y})
		}

	}

	for i := 1; i <= dis[sx][sy]; i++ {
		ans += cnt[i]
	}
	fmt.Fprintln(out, ans)
	return
}
