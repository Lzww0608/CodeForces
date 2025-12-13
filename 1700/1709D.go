package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a[i])
	}

	w := bits.Len(uint(m))
	st := make([][]int, w)
	for i := range w {
		st[i] = make([]int, m)
	}
	st[0] = a
	for i := 1; i < w; i++ {
		for j := range m - (1 << i) + 1 {
			st[i][j] = max(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}

	query := func(l, r int) int {
		k := bits.Len(uint(r-l)) - 1
		return max(st[k][l], st[k][r-(1<<k)])
	}

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		var xs, ys, xf, yf, k int
		fmt.Fscan(in, &xs, &ys, &xf, &yf, &k)
		if ys > yf {
			xs, xf = xf, xs
			ys, yf = yf, ys
		}

		if abs(xs-xf)%k != 0 || abs(ys-yf)%k != 0 {
			fmt.Fprintln(out, "NO")
			continue
		}

		if query(ys - 1, yf) < xs+(n-xs)/k*k {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
