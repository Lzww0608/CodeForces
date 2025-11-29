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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	var s string
	fmt.Fscan(in, &s)
	p := 0
	for i := range n - 1 {
		if s[i] == 'L' {
			p++
		}
	}

	l, r := p, p
	x := a[p] % m
	ans := make([]int, n)
	ans[n-1] = x
	for i := n - 2; i >= 0; i-- {
		if s[i] == 'L' {
			l--
			x = x * a[l] % m
		} else {
			r++
			x = x * a[r] % m
		}
		ans[i] = x
	}

	for i := range n {
		fmt.Fprintf(out, "%d ", ans[i])
	}
	fmt.Fprintln(out)
}
