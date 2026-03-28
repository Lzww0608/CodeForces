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
	var a, b, c string
	fmt.Fscan(in, &a, &b, &c)
	m, n := len(a), len(b)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := range a {
		if a[i] != c[i] {
			f[i+1][0] = f[i][0] + 1
		} else {
			f[i+1][0] = f[i][0]
		}
	}

	for j := range b {
		if b[j] != c[j] {
			f[0][j+1] = f[0][j] + 1
		} else {
			f[0][j+1] = f[0][j]
		}
	}

	for i := range m {
		for j := range n {
			if a[i] != c[i+j+1] {
				f[i+1][j+1] = f[i][j+1] + 1
			} else {
				f[i+1][j+1] = f[i][j+1]
			}

			if b[j] != c[i+j+1] {
				f[i+1][j+1] = min(f[i+1][j+1], f[i+1][j]+1)
			} else {
				f[i+1][j+1] = min(f[i+1][j+1], f[i+1][j])
			}
		}
	}

	fmt.Fprintln(out, f[m][n])
}
