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

	var n, m, x, k int
	type pair struct {
		s string
		x int
	}
	fmt.Fscan(in, &n, &m)
	s := make([][]pair, m)
	var t string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t, &x, &k)
		s[x-1] = append(s[x-1], pair{t, k})
	}
	for k := 0; k < m; k++ {
		sort.Slice(s[k], func(i, j int) bool {
			return s[k][i].x > s[k][j].x
		})
	}
	for i := 0; i < m; i++ {
		if len(s[i]) > 2 && s[i][1].x == s[i][2].x {
			fmt.Fprintln(out, "?")
		} else {
			fmt.Fprintln(out, s[i][0].s, s[i][1].s)
		}
	}
}
