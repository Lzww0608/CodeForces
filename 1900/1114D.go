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
	fmt.Fscan(in, &n)

	m := 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] != a[m] {
			m++
			a[m] = a[i]
		}
	}

	m++
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, m)
	}

	for d := 1; d < m; d++ {
		for l := 0; l+d < m; l++ {
			r := l + d
			if a[l] == a[r] {
				f[l][r] = f[l+1][r-1] + 1
			} else {
				f[l][r] = min(f[l+1][r], f[l][r-1]) + 1
			}
		}
	}

	fmt.Fprintln(out, f[0][m-1])
}
