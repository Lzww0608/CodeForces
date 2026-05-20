package main

import (
	"bufio"
	"fmt"
	"os"
)

const inf int = 1e9

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, l, k int
	fmt.Fscan(in, &n, &l, &k)
	d := make([]int, n+1)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &d[i])
	}
	d[n] = l
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][0] = 0

	for i := range n {
		for j := 0; j <= i; j++ {
			for x := i + 1; x <= n; x++ {
				f[x][j+1] = min(f[x][j+1], f[i][j]+(d[x]-d[i])*a[i])
			}
		}
	}

	ans := inf
	for i := n - k; i <= n; i++ {
		ans = min(ans, f[n][i])
	}
	fmt.Fprintln(out, ans)
}
