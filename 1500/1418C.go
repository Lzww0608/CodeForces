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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	f := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		f[i] = [2]int{n, n}
	}
	f[0][1] = 0

	for i, x := range a {
		for j := 0; j < 2; j++ {
			f[i+1][j^1] = min(f[i+1][j^1], f[i][j]+j*x)
			if i+2 <= n {
				f[i+2][j^1] = min(f[i+2][j^1], f[i][j]+j*(x+a[i+1]))
			}
		}
	}
	fmt.Fprintln(out, min(f[n][0], f[n][1]))
}
