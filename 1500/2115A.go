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
	a := make([]int, n)
	g := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		g = gcd(a[i], g)
	}

	cnt := 0
	for _, x := range a {
		if x == g {
			cnt++
		}
	}

	if cnt != 0 {
		fmt.Fprintln(out, n-cnt)
		return
	}

	mx := slices.Max(a)
	f := make([]int, mx+1)
	for i := range f {
		f[i] = n + 1
	}
	for _, x := range a {
		x /= g
		g := make([]int, mx+1)
		copy(g, f)
		g[x] = 1

		for i := 1; i <= mx; i++ {
			if f[i] == n+1 {
				continue
			}
			j := gcd(i, x)
			if g[j] > f[i]+1 {
				g[j] = f[i] + 1
			}
		}
		f = g
	}

	k := f[1]
	fmt.Fprintln(out, n+k-2)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
