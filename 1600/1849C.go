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
	var s string
	fmt.Fscan(in, &s)

	f := make([]int, n)
	g := make([]int, n)
	for i := range s {
		f[i] = n
		if i == 0 {
			if s[i] == '0' {
				f[i] = i
			}
		} else {
			if s[i] == '0' {
				f[i] = i
			} else {
				f[i] = f[i-1]
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		g[i] = -1
		if i == n-1 {
			if s[i] == '1' {
				g[i] = i
			}
		} else {
			if s[i] == '1' {
				g[i] = i
			} else {
				g[i] = g[i+1]
			}
		}
	}

	vis := make(map[pair]bool)
	ans := 0

	var l, r int
	for range m {
		fmt.Fscan(in, &l, &r)
		l--
		r--

		x, y := g[l], f[r]
		if x > y || x == -1 || y == n {
			ans |= 1
		} else {
			vis[pair{x, y}] = true
		}
	}

	fmt.Fprintln(out, ans+len(vis))
}

type pair struct {
	x, y int
}
