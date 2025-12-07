package main

import (
	"bufio"
	"fmt"
	"os"
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
	var n, k, q int
	fmt.Fscan(in, &n, &k, &q)
	a := make([]int, k+1)
	b := make([]int, k+1)
	a[0] = 0
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i+1])
	}
	b[0] = 0
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &b[i+1])
	}

	c := make([]int, q)
	id := make([]int, q)
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		id[i] = i
		fmt.Fscan(in, &c[i])
	}
	sort.Slice(id, func(i, j int) bool {
		return c[id[i]] < c[id[j]]
	})

	for i, j := 0, 0; i <= k && j < q; i++ {
		for j < q && a[i] == c[id[j]] {
			ans[id[j]] = b[i]
			j++
		}
		for j < q && i < k && a[i+1] > c[id[j]] {
			ans[id[j]] = b[i] + (b[i+1]-b[i])*(c[id[j]]-a[i])/(a[i+1]-a[i])
			j++
		}
	}

	for i := range q {
		fmt.Fprintf(out, "%d ", ans[i])
	}
	fmt.Fprintln(out)
}
