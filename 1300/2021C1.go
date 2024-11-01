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
	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)
	a := make([]int, n)
	b := make([]int, m)
	p := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
		p[a[i]] = i
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
		b[i]--
	}
	k := 0
	for _, x := range b {
		if p[x] > k {
			fmt.Fprintln(out, "TIDAK")
			return
		} else if p[x] == k {
			k++
		}

	}
	fmt.Fprintln(out, "YA")
	return
}
