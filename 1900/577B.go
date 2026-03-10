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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	if m <= n {
		fmt.Fprintln(out, "YES")
		return
	}

	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m)
	}
	for i, x := range a {
		f[i+1][x%m] = true
		for j := range f[i] {
			if f[i][j] {
				f[i+1][(j+x)%m] = true
				f[i+1][j] = true
			}
		}
	}
	if f[n][0] {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
