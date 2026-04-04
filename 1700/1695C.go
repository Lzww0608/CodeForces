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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	if (n+m)&1 == 0 {
		fmt.Fprintln(out, "NO")
		return
	}

	f := make([][2]int, n*m)
	for i := range n {
		for j, x := range a[i] {
			k := i*m + j
			p := i*m + j - 1
			q := (i-1)*m + j
			if i == 0 {
				if j == 0 {
					f[k][0], f[k][1] = x, x
				} else {
					f[k][0] = f[p][0] + x
					f[k][1] = f[p][1] + x
				}
			} else if j == 0 {
				f[k][0] = f[q][0] + x
				f[k][1] = f[q][1] + x
			} else {
				f[k][0] = min(f[p][0], f[q][0]) + x
				f[k][1] = max(f[p][1], f[q][1]) + x
			}
		}
	}

	if f[len(f)-1][0] <= 0 && f[len(f)-1][1] >= 0 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
