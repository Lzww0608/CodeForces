package main

import (
	"bufio"
	"fmt"
	"os"
)

// codeforces 1792C

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
	p := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		p[a[i]] = i + 1
	}

	k := n / 2
	for k > 0 && p[k] < p[k+1] && p[n-k] < p[n-k+1] {
		k--
	}

	fmt.Fprintln(out, k)
}
