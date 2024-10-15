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
	fa := make([]int, m*n+1)
	a := make([]int, m*n+1)
	b := make([]int, m*n+1)
	id := make([]int, m*n+1)
	R := make([]int, n+1)
	C := make([]int, m+1)

	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	for i := 1; i <= m*n; i++ {
		fmt.Fscan(in, &a[i])
		fa[i], id[i] = i, i
	}

	sort.Slice(id[1:], func(i, j int) bool {
		return a[id[i+1]] < a[id[j+1]]
	})

	for _, i := range id[1:] {
		r, c := (i-1)/m+1, (i-1)%m+1
		fr, fc := find(R[r]), find(C[c])
		p := find(i)
		x, y := b[fr], b[fc]
		if a[p] > a[fr] {
			x++
		} else {
			fa[fr] = p
		}

		if a[p] > a[fc] {
			y++
		} else {
			fa[fc] = p
		}
		b[p] = max(x, y)
		R[r], C[c] = p, p
	}
	for i := 1; i <= m*n; i++ {
		fmt.Fprint(out, b[find(i)], " ")
		if i%m == 0 {
			fmt.Fprintln(out)
		}
	}
}
