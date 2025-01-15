package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs [][]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	s := make([][]byte, n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	cnt1 := 0
	bad := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == 'G' {
				cnt1++
			} else if s[i][j] == 'B' {
				bad++
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x < 0 || x >= n || y < 0 || y >= m {
						continue
					}
					if s[x][y] == 'G' {
						fmt.Fprintln(out, "No")
						return
					} else if s[x][y] == '.' {
						s[x][y] = '#'
					}
				}
			}
		}
	}

	if cnt1 == 0 {
		fmt.Fprintln(out, "Yes")
		return
	} else if s[n-1][m-1] == '#' && bad > 0 {
		fmt.Fprintln(out, "No")
		return
	}

	cnt2 := 0
	q := [][]int{{n - 1, m - 1}}
	vis := make([]bool, m*n)
	vis[m*n-1] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		i, j := cur[0], cur[1]
		if s[i][j] == 'G' {
			cnt2++
		}
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= n || y < 0 || y >= m || s[x][y] == '#' || s[x][y] == 'B' {
				continue
			}
			if !vis[x*m+y] {
				vis[x*m+y] = true
				q = append(q, []int{x, y})
			}
		}
	}

	if cnt1 == cnt2 {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
	return
}
