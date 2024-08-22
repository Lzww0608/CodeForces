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

	var n int
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
		}
	}

	var dfs func(int, int, int) int
	dfs = func(l, r, k int) int {
		if l > r {
			return 0
		} else if l == r {
			return a[k]
		}
		if f[l][r][k] != 0 {
			return f[l][r][k]
		}
		f[l][r][k] = dfs(l, r-1, 0) + a[k]
		for i := l; i < r; i++ {
			if s[i] == s[r] {
				f[l][r][k] = max(f[l][r][k], dfs(l, i, k+1)+dfs(i+1, r-1, 0))
			}

		}
		return f[l][r][k]
	}
	fmt.Fprintln(out, dfs(0, n-1, 0))
}
