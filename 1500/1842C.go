package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// https://codeforces.com/problemset/problem/1842/C
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

	// f[i] = max(f[j] + i - j + 1)
	pos := make(map[int]int)
	f := make([]int, n+1)
	for i, x := range a {
		if v, ok := pos[x]; !ok {
			f[i+1] = f[i]
			pos[x] = f[i] - i + 1
		} else {
			f[i+1] = max(f[i], v+i)
			pos[x] = max(pos[x], f[i]-i+1)
		}
	}

	fmt.Fprintln(out, slices.Max(f))
}
