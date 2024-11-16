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
	fa1 := make([]int, n+1)
	fa2 := make([]int, n+1)
	for i := 0; i <= n; i++ {
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
		merge(fa1, x, y)
	}
	for i := 0; i < m2; i++ {
		fmt.Fscan(in, &x, &y)
		merge(fa2, x, y)
	}

	ans := [][]int{}
	for i := 1; i <= n; i++ {
		a1, b1 := find(fa1, 1), find(fa1, i)
		a2, b2 := find(fa2, 1), find(fa2, i)
		if a1 != b1 && a2 != b2 {
			ans = append(ans, []int{i, 1})
			merge(fa1, a1, b1)
			merge(fa2, a2, b2)
		}
	}
	st1 := []int{}
	st2 := []int{}
	for i := 1; i <= n; i++ {
		a1, b1 := find(fa1, 1), find(fa1, i)
		a2, b2 := find(fa2, 1), find(fa2, i)
		if a1 != b1 {
			st1 = append(st1, i)
		}
		if a2 != b2 {
			st2 = append(st2, i)
		}
	}
	for len(st1) > 0 && len(st2) > 0 {
		a, b := st1[len(st1)-1], st2[len(st2)-1]
		ans = append(ans, []int{a, b})
		merge(fa1, a, b)
		merge(fa2, a, b)
		for len(st1) > 0 && find(fa1, 1) == find(fa1, st1[len(st1)-1]) {
			st1 = st1[:len(st1)-1]
		}
		for len(st2) > 0 && find(fa2, 1) == find(fa2, st2[len(st2)-1]) {
			st2 = st2[:len(st2)-1]
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}
