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

	var n, m int
	fmt.Fscan(in, &n, &m)
	pos := make([]int, n+1)
	a := make([]int, n+m+2)
	mn := make([]int, n+1)
	mx := make([]int, n+1)

	update := func(i, v int) {
		for ; i < len(a); i += i & (-i) {
			a[i] += v
		}
	}

	query := func(i int) int {
		res := 0
		for ; i > 0; i -= i & (-i) {
			res += a[i]
		}
		return res
	}

	for i := 1; i <= n; i++ {
		pos[i] = i + m
		update(i+m, 1)
		mn[i], mx[i] = i, i
	}

	var x int
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &x)
		mn[x] = 1
		mx[x] = max(mx[x], query(pos[x]))
		update(pos[x], -1)
		pos[x] = m - i + 1
		update(m-i+1, 1)
	}

	for i := 1; i <= n; i++ {
		mx[i] = max(mx[i], query(pos[i]))
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d %d\n", mn[i], mx[i])
	}
}
