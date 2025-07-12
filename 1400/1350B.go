package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	f := make([]int, n+1)
	for i := 1; i*2 <= n; i++ {
		for j := i * 2; j <= n; j += i {
			if a[j] > a[i] {
				f[j] = max(f[j], f[i]+1)
			}
		}
	}

	fmt.Fprintln(out, slices.Max(f)+1)
}
