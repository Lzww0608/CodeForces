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

	var n, m1, m2 int
	fmt.Fscan(in, &n, &m1, &m2)
	fa1 := make([]int, n)
	fa2 := make([]int, n)
	for i := 0; i < n; i++ {
		fa1[i], fa2[i] = i, i
	}

	var find func([]int, int) int
	find = func(fa []int, x int) int {
		if fa[x] != x {
			fa[x] = find(fa, fa[x])
		}
		return fa[x]
	}

	merge := func(fa []int, x, y int) {
		rx, ry := find(fa, x), find(fa, y)
		if rx != ry {
			fa[rx] = ry
		}
	}

	var x, y int
	for i := 0; i < m1; i++ {
		fmt.Fscan(in, &x, &y)
		merge(fa1, x-1, y-1)
	}
	for i := 0; i < m2; i++ {
		fmt.Fscan(in, &x, &y)
		merge(fa2, x-1, y-1)
	}

	ans := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a1, b1 := find(fa1, i), find(fa1, j)
			a2, b2 := find(fa2, j), find(fa2, i)
			if a1 != b1 && a2 != b2 {
				ans = append(ans, []int{i, j})
				merge(fa1, a1, b1)
				merge(fa2, a2, b2)
			}
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0]+1, v[1]+1)
	}
}
