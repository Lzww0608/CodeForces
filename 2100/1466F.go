package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	fa := make([]int, m+2)
	for i := range fa {
		fa[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) bool {
		rx, ry := find(x), find(y)
		if rx != ry {
			fa[rx] = ry
			return true
		}
		return false
	}

	ans := []int{}
	res := 1
	for i := 1; i <= n; i++ {
		var k int
		fmt.Fscan(in, &k)
		a, b := -1, m+1
		fmt.Fscan(in, &a)
		if k > 1 {
			fmt.Fscan(in, &b)
		}
		if merge(a, b) {
			ans = append(ans, i)
			res = res * 2 % MOD
		}
	}
	fmt.Fprintln(out, res, len(ans))
	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
}
