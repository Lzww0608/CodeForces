package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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
	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	f := make([]int, n+1)
	for i, y := range a {
		cnt := (x + y - 1) / y
		if i+cnt <= n {
			f[i+cnt] = max(f[i+cnt], f[i]+1)
		}
	}

	fmt.Fprintln(out, slices.Max(f))
}
