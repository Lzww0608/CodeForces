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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	id := make([]int, n)
	ans := 0
	for i := range a {
		id[i] = i
		fmt.Fscan(in, &a[i])
		ans += a[i]
	}
	sort.Slice(id, func(i, j int) bool {
		return a[id[i]]-(n-id[i]) > a[id[j]]-(n-id[j])
	})

	vis := make(map[int]bool)
	for _, i := range id[:k] {
		vis[i] = true
		ans -= a[i]
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if vis[i] {
			cnt++
		} else {
			ans += cnt
		}
	}
	fmt.Fprintln(out, ans)
}
