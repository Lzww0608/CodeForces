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

	var n, m int
	fmt.Fscan(in, &n, &m)
	b := make([]int, n)
	g := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &g[i])
	}
	sort.Ints(b)
	sort.Ints(g)
	if g[0] < b[n-1] {
		fmt.Fprintln(out, -1)
		return
	}
	ans := 0
	for i := range b {
		ans += b[i] * m
	}
	for i := 1; i < m; i++ {
		ans += g[i] - b[n-1]
	}
	if g[0] == b[n-1] {
		ans += g[0] - b[n-1]
	} else {
		ans += g[0] - b[n-2]
	}
	fmt.Fprintln(out, ans)
}
