package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// codeforces 1324E
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, h, l, r int
	fmt.Fscan(in, &n, &h, &l, &r)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, h)
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}
	f[0][0] = 0

	for i := range a {
		for j := range h {
			x := (j - a[i] + h) % h
			y := (j - a[i] + 1 + h) % h
			f[i+1][j] = max(f[i][x], f[i][y])
			if j >= l && j <= r {
				f[i+1][j]++
			}
		}
	}

	ans := 0
	for j := range h {
		ans = max(ans, f[n][j])
	}
	fmt.Fprintln(out, ans)
}
