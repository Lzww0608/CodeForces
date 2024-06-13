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

	var t, n, m, x1, y1, x2, y2 int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &m)
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		a, b := 4, 4
		dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range dirs {
			if x1+dir[0] > n || x1+dir[0] <= 0 || y1+dir[1] > m || y1+dir[1] <= 0 {
				a--
			}
			if x2+dir[0] > n || x2+dir[0] <= 0 || y2+dir[1] > m || y2+dir[1] <= 0 {
				b--
			}
		}
		fmt.Fprintln(out, min(a, b))
	}
}
