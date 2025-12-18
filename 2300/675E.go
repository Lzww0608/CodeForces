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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	a[n-1] = n - 1

	cmp := func(i, j int) int {
		if a[i] < a[j] {
			return j
		}
		return i
	}

	w := bits.Len(uint(n))
	st := make([][]int, w)
	for i := range w {
		st[i] = make([]int, n)
	}
	for i := range n {
		st[0][i] = i
	}
	for i := 1; i < w; i++ {
		for j := range n - (1 << i) + 1 {
			st[i][j] = cmp(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}

	query := func(l, r int) int {
		k := bits.Len(uint(r-l)) - 1
		return cmp(st[k][l], st[k][r-(1<<k)])
	}

	f := make([]int, n)
	ans := 0
	for i := n - 2; i >= 0; i-- {
		j := query(i+1, a[i]+1)
		f[i] = f[j] + n - 1 - i - a[i] + j
		ans += f[i]
	}

	fmt.Fprintln(out, ans)
}
