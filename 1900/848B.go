package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 100_005

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, w, h int
	fmt.Fscan(in, &n, &w, &h)
	s := make([][]int, N*2)
	g := make([]int, N)
	p := make([]int, N)
	t := make([]int, N)
	ans_x := make([]int, n)
	ans_y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &g[i], &p[i], &t[i])
		s[p[i]-t[i]+N] = append(s[p[i]-t[i]+N], i)
	}

	for i := 0; i < N*2; i++ {
		if len(s[i]) == 0 {
			continue
		}

		xs := []int{}
		ys := []int{}

		for _, u := range s[i] {
			if g[u] == 1 {
				xs = append(xs, p[u])
			} else {
				ys = append(ys, p[u])
			}
		}

		sort.Ints(xs)
		sort.Ints(ys)
		sort.Slice(s[i], func(u, v int) bool {
			if g[s[i][u]] != g[s[i][v]] {
				return g[s[i][u]] == 2
			}
			if g[s[i][u]] == 2 {
				return p[s[i][u]] > p[s[i][v]]
			}
			return p[s[i][u]] < p[s[i][v]]
		})

		for j, x := range xs {
			ans_x[s[i][j]] = x
			ans_y[s[i][j]] = h
		}

		for j := range ys {
			ans_x[s[i][j+len(xs)]] = w
			ans_y[s[i][j+len(xs)]] = ys[len(ys)-j-1]
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ans_x[i], ans_y[i])
	}
}
