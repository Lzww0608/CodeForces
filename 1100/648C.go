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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	q := [][]int{}
	vis := make([]bool, m*n)
	for i := 0; i < n && len(q) == 0; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == 'S' {
				q = append(q, []int{i, j, -1, -1})
				break
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		i, j := node[0], node[1]
		for k, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < n && y >= 0 && y < m && (s[x][y] == '*' || s[x][y] == 'S') && !vis[x*m+y] && (node[2] != x || node[3] != y) {
				vis[x*m+y] = true
				q = append(q, []int{x, y, i, j})
				switch k {
				case 0:
					fmt.Fprint(out, "D")
				case 1:
					fmt.Fprint(out, "U")
				case 2:
					fmt.Fprint(out, "R")
				case 3:
					fmt.Fprint(out, "L")
				}
				break
			}
		}
	}
	fmt.Fprintln(out)
}
