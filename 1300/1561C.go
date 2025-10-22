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
	var n int
	fmt.Fscan(in, &n)

	a := make([][]int, n)
	b := make([]int, n)
	id := make([]int, n)
	for i := range n {
		var k int
		fmt.Fscan(in, &k)
		a[i] = make([]int, k)
		mx := 0
		for j := range k {
			fmt.Fscan(in, &a[i][j])
			mx = max(mx, a[i][j]-j)
		}
		b[i] = mx
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return b[id[i]] < b[id[j]]
	})

	cur, ans := 0, 1
	for _, i := range id {
		k := len(a[i])
		x := b[i]
		if cur < x {
			ans += x - cur
			cur = x
		}
		cur += k
	}

	fmt.Fprintln(out, ans)
}
